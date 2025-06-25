package set1

import (
	"errors"
)

func XorFixedLength(string1 string, string2 string) (string, error) {
	if len(string1) != len(string2) {
		return "", errors.New("two input strings have different lengths")
	}

	bytes1, err := HexToBytes(string1)
	if err != nil {
		return "", err
	}
	bytes2, err := HexToBytes(string2)
	if err != nil {
		return "", err
	}

	var result []byte
	for i := 0; i < len(bytes1); i++ {
		result = append(result, bytes1[i]^bytes2[i])
	}

	return BytesToHex(result), nil
}
