package set1

import (
	"errors"
	"strings"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func HexToBase64(hexInput string) string {

	bytes, err := toByteSlice(hexInput)
	if err != nil {
		println(err.Error())
	}
	//fmt.Printf("Byte (decimal): %v\n", bytes)
	//fmt.Printf("Byte (hex):     0x%02x\n", bytes)

	//fmt.Println("grouped into 6 bits")
	groupOfSixBits := groupBySixBits(bytes)
	//for _, val := range groupOfSixBits {
	//	fmt.Printf("%v\n", val)
	//}
	//fmt.Println("grouped into 6 bits")

	output := toBase64Str(groupOfSixBits)
	//fmt.Printf("base64 output: %s\n", output)

	return output
}

// from hex string to byte slice
// e.g. hex: "12", output: 1*16 + 2 = 18 as the first byte [0x12]
func toByteSlice(s string) ([]byte, error) {
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

// group the bytes into unit of 6, each 6 bits group is stored in a uint8
func groupBySixBits(bytes []byte) []uint8 {
	var result []uint8
	var buffer uint32 = 0
	var bitsInBuffer uint8 = 0

	for _, b := range bytes {
		buffer = (buffer << 8) | uint32(b)
		bitsInBuffer += 8

		for bitsInBuffer >= 6 {
			shift := bitsInBuffer - 6
			group := (uint8(buffer >> shift)) & 0b111111
			result = append(result, group)
			bitsInBuffer -= 6
		}
	}

	if bitsInBuffer > 0 {
		buffer <<= 6 - bitsInBuffer
		result = append(result, uint8(buffer)&0b111111)
	}

	return result
}

func toBase64Str(sixBitGroups []uint8) string {
	var res strings.Builder
	for _, val := range sixBitGroups {
		base64Char := base64Table[val]
		res.WriteByte(base64Char)
	}

	mod := len(res.String()) % 4
	if mod > 0 {
		for i := 0; i < 4-mod; i++ {
			res.WriteByte('=')
		}
	}
	return res.String()
}
