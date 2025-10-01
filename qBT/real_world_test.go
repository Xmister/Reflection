package qBT

import (
	"encoding/json"
	"testing"
)

// TestRealWorldPayload tests with a payload similar to what was saved in the error
func TestRealWorldPayload(t *testing.T) {
	// This is based on the actual payload from the error message but simplified
	realWorldJSON := `{
		"add_stopped_enabled": false,
		"add_to_top_of_queue": false,
		"proxy_type": "None",
		"proxy_ip": "",
		"proxy_port": 8080,
		"proxy_auth_enabled": false,
		"proxy_username": "",
		"proxy_password": "",
		"save_path": "/mnt/sata/Torrent",
		"web_ui_port": 8088
	}`

	var pref Preferences
	err := json.Unmarshal([]byte(realWorldJSON), &pref)
	if err != nil {
		t.Fatalf("Failed to unmarshal real-world JSON payload: %v", err)
	}

	// Verify the proxy_type was correctly parsed
	if pref.Proxy_type != ProxyTypeNone {
		t.Errorf("Expected proxy_type to be ProxyTypeNone (0), got %d", pref.Proxy_type)
	}

	// Verify other fields are correct
	if pref.Proxy_ip != "" {
		t.Errorf("Expected proxy_ip to be empty, got %s", pref.Proxy_ip)
	}

	if pref.Proxy_port != 8080 {
		t.Errorf("Expected proxy_port to be 8080, got %d", pref.Proxy_port)
	}

	if pref.Save_path != "/mnt/sata/Torrent" {
		t.Errorf("Expected save_path to be '/mnt/sata/Torrent', got %s", pref.Save_path)
	}

	if pref.Web_ui_port != 8088 {
		t.Errorf("Expected web_ui_port to be 8088, got %d", pref.Web_ui_port)
	}
}

// TestMixedProxyTypes tests different proxy types in the same JSON
func TestMixedProxyTypes(t *testing.T) {
	testCases := []struct {
		name     string
		json     string
		expected ProxyType
	}{
		{
			name:     "None proxy",
			json:     `{"proxy_type": "None"}`,
			expected: ProxyTypeNone,
		},
		{
			name:     "HTTP proxy",
			json:     `{"proxy_type": "HTTP"}`,
			expected: ProxyTypeHTTP,
		},
		{
			name:     "SOCKS5 proxy",
			json:     `{"proxy_type": "SOCKS5"}`,
			expected: ProxyTypeSOCKS5,
		},
		{
			name:     "SOCKS4 proxy",
			json:     `{"proxy_type": "SOCKS4"}`,
			expected: ProxyTypeSOCKS4,
		},
		{
			name:     "Legacy integer 0",
			json:     `{"proxy_type": 0}`,
			expected: ProxyTypeNone,
		},
		{
			name:     "Legacy integer 1",
			json:     `{"proxy_type": 1}`,
			expected: ProxyTypeHTTP,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var pref Preferences
			err := json.Unmarshal([]byte(tc.json), &pref)
			if err != nil {
				t.Fatalf("Failed to unmarshal JSON for %s: %v", tc.name, err)
			}

			if pref.Proxy_type != tc.expected {
				t.Errorf("For %s: expected proxy_type %d, got %d", tc.name, tc.expected, pref.Proxy_type)
			}
		})
	}
}