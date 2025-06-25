package set1

import "testing"

func TestXorFixedLength(t *testing.T) {
	tests := []struct {
		hexInput1 string
		hexInput2 string
		expected  string
	}{
		{"1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965", "746865206b696420646f6e277420706c6179"},
	}

	for _, tt := range tests {
		result, _ := XorFixedLength(tt.hexInput1, tt.hexInput2)
		if result != tt.expected {
			t.Errorf("XorFixedLength(%q, %q) = %q; want %q", tt.hexInput1, tt.hexInput2, result, tt.expected)
		}
	}
}
