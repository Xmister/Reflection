package main

import (
	"testing"
	"net/http"
	"gopkg.in/h2non/gock.v1"
	log "github.com/sirupsen/logrus"
)

func TestSeparateConnections(t *testing.T) {
	const apiAddr = "http://localhost:8080"
	log.SetLevel(log.ErrorLevel) // Reduce noise

	defer gock.Off()

	// Mock login endpoint to be called multiple times
	gock.New(apiAddr).
		Post("/api/v2/auth/login").
		Times(10).
		Reply(200).
		SetHeader("Set-Cookie", "SID=test123")

	// Mock other endpoints
	gock.New(apiAddr).
		Get("/api/v2/app/version").
		Times(10).
		Reply(200).
		BodyString("4.1.0")

	gock.New(apiAddr).
		Get("/api/v2/app/preferences").
		Times(10).
		Reply(200).
		File("testdata/preferences.json")

	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	// Initialize connection factory parameters
	qbtAPIAddr = apiAddr
	qbtHTTPClient = client
	qbtUseSync = false

	// Create multiple connections to verify they're separate
	conn1 := createQBTConnection()
	conn2 := createQBTConnection()
	
	// Verify they are different instances
	if conn1 == conn2 {
		t.Error("Expected separate connection instances, but got the same instance")
	}
	
	// Verify they both can log in independently
	success1 := conn1.Login("test", "test")
	success2 := conn2.Login("test", "test")
	
	if !success1 {
		t.Error("First connection login failed")
	}
	
	if !success2 {
		t.Error("Second connection login failed")
	}
	
	// Verify they maintain separate auth state
	if !conn1.IsLoggedIn() {
		t.Error("First connection should be logged in")
	}
	
	if !conn2.IsLoggedIn() {
		t.Error("Second connection should be logged in")
	}
	
	t.Log("SUCCESS: Separate connections created and working independently")
}