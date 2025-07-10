package set1

import "testing"

func TestDetect(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc", "cc"},
	}

	for range tests {
		_, _ = Detect()
	}
}
