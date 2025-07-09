package set1

import (
	"fmt"
	"math"
	"unicode"
)

// Decipher the encrypted text (HEX) that is single byte xor'd
// encrypted is the single-byte xor'd text
// key is the xor character, if this is nil, this function will
// return a guessed decrypted string based on Englishness (letter freq)
func Decipher(encrypted string, key *byte) (string, error) {

	inputInBytes, err := HexToBytes(encrypted)
	if err != nil {
		return "", err
	}

	if key != nil {
		xordBytes := make([]byte, len(inputInBytes))
		for idx, char := range inputInBytes {
			xordBytes[idx] = char ^ *key
		}
		return string(xordBytes), nil
	}

	var result string
	var minLoss = math.MaxFloat64
	var freqLoss float64

	// trying out each char from ascii as key
	for i := 0; i <= 255; i++ {
		guessKey := byte(i)

		xordBytes := make([]byte, len(inputInBytes))
		for idx, char := range inputInBytes {
			xordBytes[idx] = char ^ guessKey
		}

		freqLoss = getDecipheredFreqLoss(string(xordBytes))

		if freqLoss < minLoss {
			minLoss = freqLoss
			result = string(xordBytes)
			fmt.Printf("with key %s, the deciphered is %s, loss is %f\n", string(guessKey), string(xordBytes), freqLoss)
		}
	}

	return result, nil
}

func getDecipheredFreqLoss(deciphered string) float64 {
	loss := 0.0
	letterFreqMap := getLetterFreqMap(deciphered)

	for l, letterFreq := range letterFreqMap {
		englishFreq := getStandardEnglishFrequency(l)
		loss += math.Abs(letterFreq - englishFreq)
	}

	englishCommonChar := getEnglishCommonChar()
	for _, l := range deciphered {
		if !unicode.IsLetter(l) && !englishCommonChar[l] {
			loss += 50.0 // giving a penalty for uncommon chars
		}
	}
	return loss
}

func getLetterFreqMap(text string) map[rune]float64 {
	letterCountMap := make(map[rune]int)
	var totalLetterCount int
	for _, char := range text {
		if unicode.IsLetter(char) {
			letterCountMap[unicode.ToLower(char)]++
			totalLetterCount++
		}
	}
	letterFreqMap := make(map[rune]float64)
	for k, v := range letterCountMap {
		letterFreqMap[k] = float64(v / totalLetterCount)
	}
	return letterFreqMap
}

func getStandardEnglishFrequency(character rune) float64 {
	// Source: https://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html
	lowerCased := unicode.ToLower(character)
	var englishLetterFrequencies = map[rune]float64{
		'a': 8.167,
		'b': 1.492,
		'c': 2.782,
		'd': 4.253,
		'e': 12.702,
		'f': 2.228,
		'g': 2.015,
		'h': 6.094,
		'i': 6.966,
		'j': 0.153,
		'k': 0.772,
		'l': 4.025,
		'm': 2.406,
		'n': 6.749,
		'o': 7.507,
		'p': 1.929,
		'q': 0.095,
		'r': 5.987,
		's': 6.327,
		't': 9.056,
		'u': 2.758,
		'v': 0.978,
		'w': 2.360,
		'x': 0.150,
		'y': 1.974,
		'z': 0.074,
	}

	if freq, ok := englishLetterFrequencies[lowerCased]; ok {
		return freq
	}
	return 0
}

func getEnglishCommonChar() map[rune]bool {
	return map[rune]bool{
		',': true,
		'-': true,
		'.': true,
		'!': true,
		'?': true,
		' ': true,
	}
}
