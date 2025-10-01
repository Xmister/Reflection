package qBT

import (
	"encoding/json"
	"testing"
)

func TestProxyTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		jsonInput   string
		expected    ProxyType
		expectError bool
	}{
		{
			name:      "String None",
			jsonInput: `{"proxy_type":"None"}`,
			expected:  ProxyTypeNone,
		},
		{
			name:      "String HTTP",
			jsonInput: `{"proxy_type":"HTTP"}`,
			expected:  ProxyTypeHTTP,
		},
		{
			name:      "String SOCKS5",
			jsonInput: `{"proxy_type":"SOCKS5"}`,
			expected:  ProxyTypeSOCKS5,
		},
		{
			name:      "String SOCKS4",
			jsonInput: `{"proxy_type":"SOCKS4"}`,
			expected:  ProxyTypeSOCKS4,
		},
		{
			name:      "Integer 0",
			jsonInput: `{"proxy_type":0}`,
			expected:  ProxyTypeNone,
		},
		{
			name:      "Integer 1",
			jsonInput: `{"proxy_type":1}`,
			expected:  ProxyTypeHTTP,
		},
		{
			name:      "Integer 2",
			jsonInput: `{"proxy_type":2}`,
			expected:  ProxyTypeSOCKS5,
		},
		{
			name:      "Integer 3",
			jsonInput: `{"proxy_type":3}`,
			expected:  ProxyTypeSOCKS4,
		},
		{
			name:      "String number 0",
			jsonInput: `{"proxy_type":"0"}`,
			expected:  ProxyTypeNone,
		},
		{
			name:      "String number 1",
			jsonInput: `{"proxy_type":"1"}`,
			expected:  ProxyTypeHTTP,
		},
		{
			name:        "Invalid string",
			jsonInput:   `{"proxy_type":"Invalid"}`,
			expectError: true,
		},
		{
			name:      "Null value (should use zero value)",
			jsonInput: `{"proxy_type":null}`,
			expected:  ProxyTypeNone, // Zero value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pref Preferences
			err := json.Unmarshal([]byte(tt.jsonInput), &pref)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for input %s, but got none", tt.jsonInput)
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", tt.jsonInput, err)
				return
			}
			
			if pref.Proxy_type != tt.expected {
				t.Errorf("Expected proxy_type %d, got %d for input %s", 
					tt.expected, pref.Proxy_type, tt.jsonInput)
			}
		})
	}
}

// Test the exact case from the error message
func TestProxyTypeErrorCase(t *testing.T) {
	// This is the exact JSON structure that was causing the panic
	problemJSON := `{"proxy_type":"None","proxy_ip":"","proxy_port":8080}`
	
	var pref Preferences
	err := json.Unmarshal([]byte(problemJSON), &pref)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON that was causing panic: %v", err)
	}
	
	if pref.Proxy_type != ProxyTypeNone {
		t.Errorf("Expected proxy_type to be ProxyTypeNone (0), got %d", pref.Proxy_type)
	}
	
	if pref.Proxy_ip != "" {
		t.Errorf("Expected proxy_ip to be empty string, got %s", pref.Proxy_ip)
	}
	
	if pref.Proxy_port != 8080 {
		t.Errorf("Expected proxy_port to be 8080, got %d", pref.Proxy_port)
	}
}

func TestProxyTypeString(t *testing.T) {
	tests := []struct {
		proxyType ProxyType
		expected  string
	}{
		{ProxyTypeNone, "None"},
		{ProxyTypeHTTP, "HTTP"},
		{ProxyTypeSOCKS5, "SOCKS5"},
		{ProxyTypeSOCKS4, "SOCKS4"},
		{ProxyType(99), "Unknown(99)"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.proxyType.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}