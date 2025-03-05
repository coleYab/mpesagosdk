package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConstructURL(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		endpoint string
		expected string
	}{
		{
			name:     "Production Environment",
			env:      "PRODUCTION",
			endpoint: "/v1/payments",
			expected: "https://api.safaricom.et/v1/payments",
		},
		{
			name:     "Sandbox Environment",
			env:      "SANDBOX",
			endpoint: "/v1/payments",
			expected: "https://apisandbox.safaricom.et/v1/payments",
		},
		{
			name:     "Invalid Environment",
			env:      "INVALID_ENV",
			endpoint: "/v1/payments",
			expected: "https://apisandbox.safaricom.et/v1/payments", // Default to Sandbox
		},
		{
			name:     "Empty Endpoint",
			env:      "PRODUCTION",
			endpoint: "",
			expected: "https://api.safaricom.et",
		},
		{
			name:     "Endpoint with Slash",
			env:      "SANDBOX",
			endpoint: "/v1/payments/",
			expected: "https://apisandbox.safaricom.et/v1/payments/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConstructURL(tt.env, tt.endpoint)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBaseUrl(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		expected string
	}{
		{
			name:     "Production Environment",
			env:      "PRODUCTION",
			expected: "https://api.safaricom.et",
		},
		{
			name:     "Sandbox Environment",
			env:      "SANDBOX",
			expected: "https://apisandbox.safaricom.et",
		},
		{
			name:     "Invalid Environment",
			env:      "INVALID_ENV",
			expected: "https://apisandbox.safaricom.et", // Default to Sandbox
		},
		{
			name:     "Empty Environment",
			env:      "",
			expected: "https://apisandbox.safaricom.et", // Default to Sandbox
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := baseUrl(tt.env)
			assert.Equal(t, tt.expected, result)
		})
	}
}
