package set1

import (
	"errors"
	"strings"
)

/**
 * Helper class for translating between hexadecimal and byte slice
 */

// HexToBytes from hex string to byte slice
// e.g. hex: "12", output: 1*16 + 2 = 18 as the first byte [0x12]
func HexToBytes(s string) ([]byte, error) {
	if len(s)%2 != 0 {
		return nil, errors.New("hex string length must be even")
	}

	res := make([]byte, len(s)/2)
	for i := 0; i < len(s); i += 2 {
		c1 := s[i]
		c2 := s[i+1]
		c1Val := hexCharToValue(c1)
		c2Val := hexCharToValue(c2)

		if c1Val == -1 || c2Val == -1 {
			return nil, errors.New("invalid hex string")
		}

		// can't use res = append(res, byte(c1Val*16+c2Val))
		// because it'll append to the end after
		res[i/2] = byte(c1Val*16 + c2Val)
	}

	return res, nil
}

func hexCharToValue(b byte) int {
	switch {
	case b >= '0' && b <= '9':
		return int(b - '0')
	case b >= 'A' && b <= 'F':
		return int(b - 'A' + 10)
	case b >= 'a' && b <= 'f':
		return int(b - 'a' + 10)
	default:
		return -1 // not a valid hex character
	}
}

func BytesToHex(bytes []byte) string {
	hexTable := "0123456789abcdef"
	var res strings.Builder
	for _, b := range bytes {
		leftHalf := (b & 0xF0) >> 4
		rightHalf := b & 0b1111

		res.WriteByte(hexTable[leftHalf])
		res.WriteByte(hexTable[rightHalf])
	}
	return res.String()
}
