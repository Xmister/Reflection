package main

import (
	"testing"
	"github.com/hekmon/transmissionrpc"
	"net/http/httptest"
	"net"
	"net/http"
	"gopkg.in/h2non/gock.v1"
	log "github.com/sirupsen/logrus"
)

func TestDoubleConnection(t *testing.T) {
	const apiAddr = "http://localhost:8080"
	log.SetLevel(log.DebugLevel)

	defer gock.Off()

	gock.Observe(gock.DumpRequest)

	// Basic auth and version mocks
	gock.New(apiAddr).
		Post("/api/v2/auth/login").
		Reply(200).
		SetHeader("Set-Cookie", "SID=1")

	gock.New(apiAddr).
		Get("/api/v2/app/version").
		Reply(200).
		BodyString("4.1.0")

	gock.New(apiAddr).
		Get("/api/v2/app/preferences").
		Reply(200).
		File("testdata/preferences.json")

	gock.New(apiAddr).
		Get("/api/v2/transfer/info").
		Reply(200).
		File("testdata/transfer_info.json")

	gock.New(apiAddr).
		Get("/api/v2/torrents/info").
		Reply(200).
		File("testdata/torrent_list.json")

	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	qBTConn.Init(apiAddr, client, false)
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	defer server.CloseClientConnections()
	serverAddr := server.Listener.Addr().(*net.TCPAddr)

	// First connection
	t.Log("Creating first transmission client...")
	transmissionbt1, err := transmissionrpc.New(serverAddr.IP.String(), "", "",
		&transmissionrpc.AdvancedConfig{Port: uint16(serverAddr.Port)})
	if err != nil {
		t.Fatalf("Failed to create first client: %v", err)
	}
	t.Log("First client created successfully")

	// Test first client
	session1, err := transmissionbt1.SessionArgumentsGet()
	if err != nil {
		t.Fatalf("First client session failed: %v", err)
	}
	t.Logf("First client session successful: version=%s", *session1.Version)

	// Second connection - this is where the crash might occur
	t.Log("Creating second transmission client...")
	transmissionbt2, err := transmissionrpc.New(serverAddr.IP.String(), "", "",
		&transmissionrpc.AdvancedConfig{Port: uint16(serverAddr.Port)})
	if err != nil {
		t.Fatalf("Failed to create second client: %v", err)
	}
	t.Log("Second client created successfully")

	// Test second client
	session2, err := transmissionbt2.SessionArgumentsGet()
	if err != nil {
		t.Fatalf("Second client session failed: %v", err)
	}
	t.Logf("Second client session successful: version=%s", *session2.Version)

	// Test both clients working together
	t.Log("Testing both clients again...")
	_, err = transmissionbt1.SessionArgumentsGet()
	if err != nil {
		t.Fatalf("First client second call failed: %v", err)
	}
	
	_, err = transmissionbt2.SessionArgumentsGet()
	if err != nil {
		t.Fatalf("Second client second call failed: %v", err)
	}

	t.Log("All tests passed!")
}