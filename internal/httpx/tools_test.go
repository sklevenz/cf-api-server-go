package httpx_test

import (
	"testing"

	"github.com/sklevenz/cf-api-server/internal/httpx"
)

// TestGenerateETag verifies the correct ETag is generated for known input.
func TestGenerateETag(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty data",
			input:    []byte{},
			expected: `"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"`,
		},
		{
			name:     "simple string",
			input:    []byte("hello world"),
			expected: `"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"`,
		},
		{
			name:     "longer input",
			input:    []byte("this is a longer test input for the ETag generation"),
			expected: `"54fa2f9859fca0a3da6c1d33baefa9f00defa12f474aff81e34a86f39d71499a"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := httpx.GenerateETag(tt.input)
			if result != tt.expected {
				t.Errorf("GenerateETag(%q) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}
