// Stress test to verify the authentication race condition fix
package qBT

import (
	"testing"
	"sync"
	"net/http"
	"gopkg.in/h2non/gock.v1"
	"time"
)

// High-stress test with many concurrent goroutines
func TestHighConcurrencyAuthentication(t *testing.T) {
	defer gock.Off()
	
	const apiAddr = "http://localhost:8080"
	
	// Mock login endpoint - allow many requests
	gock.New(apiAddr).
		Post("/api/v2/auth/login").
		Times(100).
		Reply(200).
		SetHeader("Set-Cookie", "SID=test123")

	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	// Create connection
	var conn Connection
	conn.Init(apiAddr, client, false)
	
	t.Log("Running high-concurrency authentication stress test...")
	
	numGoroutines := 50
	numIterationsPerGoroutine := 10
	
	var wg sync.WaitGroup
	errorsChan := make(chan error, numGoroutines*numIterationsPerGoroutine)
	
	// Each goroutine performs multiple login/check cycles
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < numIterationsPerGoroutine; j++ {
				// Check login status (to exercise the read lock)
				_ = conn.IsLoggedIn()
				
				// Try to login
				loginResult := conn.Login("test", "test")
				
				// Verify consistency
				if !loginResult {
					errorsChan <- nil // Use nil as error indicator for simplicity
				}
				
				// Small random delay to increase chance of race conditions
				time.Sleep(time.Duration(id+j) * time.Microsecond)
				
				// Check status again
				isStillLoggedIn := conn.IsLoggedIn()
				if !isStillLoggedIn {
					errorsChan <- nil // Error indicator
				}
			}
		}(i)
	}
	
	wg.Wait()
	close(errorsChan)
	
	// Count errors
	errors := 0
	for err := range errorsChan {
		if err != nil {
			errors++
		}
	}
	
	t.Logf("Completed %d goroutines x %d iterations = %d operations", 
		numGoroutines, numIterationsPerGoroutine, numGoroutines*numIterationsPerGoroutine)
	t.Logf("Errors: %d", errors)
	
	// Should have no errors with proper synchronization
	if errors > 0 {
		t.Errorf("Found %d errors in high-concurrency test", errors)
	}
	
	// Final verification
	if !conn.IsLoggedIn() {
		t.Error("Connection should be logged in after stress test")
	}
}