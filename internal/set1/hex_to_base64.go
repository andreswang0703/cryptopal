package set1

import (
	"strings"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func HexToBase64(hexInput string) string {

	bytes, err := HexToBytes(hexInput)
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
