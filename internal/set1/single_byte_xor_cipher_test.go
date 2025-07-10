package set1

import "testing"

func TestDecipher(t *testing.T) {
	tests := []struct {
		hexInput string
		expected string
	}{
		{"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736", "Cooking MC's like a pound of bacon"},
		{"75494401505448424a0143534e564f01474e59014b544c5152014e57445301554944014d405b5801454e460f", "The quick brown fox jumps over the lazy dog."},
		{"625e5f45165f451657164253454218", "This is a test."},
		{"230b440d17440544171005100d070508081d44101d1401004844070b09140d0801004408050a03110503014a", "Go is a statically typed, compiled language."},
	}

	for _, tt := range tests {
		result, _, _, err := Decipher(tt.hexInput, nil)
		if err != nil {
			t.Errorf("Decipher(%q) errored out %q; want %q", tt.hexInput, err, tt.expected)
		}
		if result != tt.expected {
			t.Errorf("HexToBase64(%q) = %q; want %q", tt.hexInput, result, tt.expected)
		}
	}
}
