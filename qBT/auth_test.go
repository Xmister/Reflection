// Test file to verify thread-safe authentication fix
package qBT

import (
	"testing"
	"sync"
	"net/http"
	"gopkg.in/h2non/gock.v1"
)

// This is a proper unit test for the authentication fix
func TestConcurrentAuthentication(t *testing.T) {
	defer gock.Off()
	
	const apiAddr = "http://localhost:8080"
	
	// Mock login endpoint
	gock.New(apiAddr).
		Post("/api/v2/auth/login").
		Times(10). // Allow multiple login attempts
		Reply(200).
		SetHeader("Set-Cookie", "SID=test123")

	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	// Create connection
	var conn Connection
	conn.Init(apiAddr, client, false)
	
	t.Log("Testing concurrent authentication...")
	
	numClients := 10
	var wg sync.WaitGroup
	results := make(chan bool, numClients)
	
	// All goroutines try to log in simultaneously
	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			success := conn.Login("test", "test")
			results <- success
		}(i)
	}
	
	wg.Wait()
	close(results)
	
	// Check results
	successes := 0
	for result := range results {
		if result {
			successes++
		}
	}
	
	t.Logf("Results: %d/%d successful logins", successes, numClients)
	
	// All should succeed due to double-check pattern
	if successes != numClients {
		t.Errorf("Expected all %d logins to succeed, but only %d succeeded", numClients, successes)
	}
	
	// Verify that IsLoggedIn works correctly
	if !conn.IsLoggedIn() {
		t.Error("Connection should show as logged in")
	}
}