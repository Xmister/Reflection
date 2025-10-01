package main

import (
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net"
	"github.com/hekmon/transmissionrpc"
	"time"
	"sync"
	"gopkg.in/h2non/gock.v1"
	log "github.com/sirupsen/logrus"
)

// Test that verifies each transmission client gets its own qBT connection
func TestTransmissionClientSeparateConnections(t *testing.T) {
	const apiAddr = "http://localhost:8080"
	
	defer gock.Off()
	
	// Mock qBittorrent API endpoints for multiple connections
	gock.New(apiAddr).
		Post("/api/v2/auth/login").
		Times(20). // Allow many login attempts
		Reply(200).
		SetHeader("Set-Cookie", "SID=test123")
		
	gock.New(apiAddr).
		Get("/api/v2/app/version").
		Times(20).
		Reply(200).
		BodyString("4.1.0")
		
	gock.New(apiAddr).
		Get("/api/v2/app/preferences").
		Times(20).
		Reply(200).
		File("testdata/preferences.json")
		
	gock.New(apiAddr).
		Get("/api/v2/transfer/info").
		Times(20).
		Reply(200).
		File("testdata/transfer_info.json")

	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	// Initialize connection factory parameters (like main() does)
	qbtAPIAddr = apiAddr
	qbtHTTPClient = client
	qbtUseSync = false
	
	// Start test server with the reflection handler
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	defer server.CloseClientConnections()
	serverAddr := server.Listener.Addr().(*net.TCPAddr)

	t.Printf("Test server started at %s:%d\n", serverAddr.IP.String(), serverAddr.Port)

	// Test multiple concurrent transmission clients
	numClients := 5
	results := make(chan error, numClients)
	
	var wg sync.WaitGroup
	
	for i := 1; i <= numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			
			t.Printf("Client %d: Creating transmission client...\n", clientID)
			
			client, err := transmissionrpc.New(serverAddr.IP.String(), "", "",
				&transmissionrpc.AdvancedConfig{Port: uint16(serverAddr.Port)})
			if err != nil {
				results <- fmt.Errorf("Client %d: Failed to create client: %v", clientID, err)
				return
			}
			
			t.Printf("Client %d: Making session request...\n", clientID)
			
			session, err := client.SessionArgumentsGet()
			if err != nil {
				results <- fmt.Errorf("Client %d: Session request failed: %v", clientID, err)
				return
			}
			
			t.Printf("Client %d: Session successful, version=%s\n", clientID, *session.Version)
			results <- nil // Success
		}(i)
		
		// Small delay to stagger requests
		time.Sleep(10 * time.Millisecond)
	}
	
	// Wait for all clients to complete
	wg.Wait()
	close(results)
	
	// Check results
	succeeded := 0
	failed := 0
	
	for err := range results {
		if err != nil {
			t.Logf("ERROR: %v", err)
			failed++
		} else {
			succeeded++
		}
	}
	
	t.Printf("\nTest Results: %d succeeded, %d failed\n", succeeded, failed)
	
	if failed > 0 {
		t.Errorf("FAILURE: %d clients failed", failed)
	} else {
		t.Logf("SUCCESS: All clients succeeded with separate connections")
	}
}