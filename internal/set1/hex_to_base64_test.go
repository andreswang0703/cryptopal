package set1

import "testing"

func TestHexToBase64(t *testing.T) {
	tests := []struct {
		hexInput string
		expected string
	}{
		{"12", "Eg=="},
		{"48656c6c6f", "SGVsbG8="}, // "Hello"
		{"4d616e", "TWFu"},         // "Man"
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
		{"", ""}, // Empty input
	}

	for _, tt := range tests {
		result := HexToBase64(tt.hexInput)
		if result != tt.expected {
			t.Errorf("HexToBase64(%q) = %q; want %q", tt.hexInput, result, tt.expected)
		}
	}
}
