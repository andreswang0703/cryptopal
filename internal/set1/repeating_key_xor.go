package set1

import (
	"encoding/hex"
)

func XorRepeatingKey(input string, key string) string {
	output := make([]byte, len(input))
	for i, c := range input {
		keyRune := key[i%len(key)]
		output[i] = byte(c) ^ keyRune
	}
	return hex.EncodeToString(output)
}
