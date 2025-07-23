package set1

import (
	"encoding/base64"
	"fmt"
	"log"
	"sort"
)

func BreakRepeatingKeyXor(data string) []string {
	// using golang base64 decoding for now, i'd want to implement my own though
	base64Decoded, _ := base64.StdEncoding.DecodeString(data)
	input := string(base64Decoded)
	//fmt.Printf("base64 decoded input: %s\n", input)

	limit := 5 // trying out 5 keySize candidates
	candidateKeySizes := findKeySizes(input, limit)
	fmt.Printf("key size: %d\n", candidateKeySizes)

	candidateOutputs := make([]string, limit)
	for i := 0; i < len(candidateKeySizes); i++ {
		keySize := candidateKeySizes[i]
		chunks := chunkBytes([]byte(input), keySize)
		transposed := transpose(chunks, keySize)

		key := calculateKeyUsingSingleXor(keySize, transposed)
		fmt.Printf("decryption key is: %s\n", string(key))

		output := XorRepeatingKey(input, string(key))
		fmt.Printf("output: %s\n", output)

		outputInBytes, err := HexToBytes(output)
		if err != nil {
			log.Fatal("failed to translate hex to bytes: ", err)
		}

		candidateOutputs[i] = string(outputInBytes)
	}

	return candidateOutputs
}

func calculateKeyUsingSingleXor(keySize int, transposed [][]byte) []rune {
	// we can use single byte xor to decipher the transposed bytes
	// each byte chunk uses same byte for xor
	key := make([]rune, keySize)
	for i := 0; i < keySize; i++ {
		hex := BytesToHex(transposed[i])
		_, _, keyRune, err := Decipher(hex, nil)
		if err != nil {
			_ = fmt.Errorf("failed to decipher single byte XOR at index %q with key size %q", i, keySize)
		}
		key[i] = keyRune
	}
	return key
}

// a struct that's used to capture the avgHummingDistance for each size attempted
// the lower, the better
type keyScore struct {
	size  int
	score float32
}

// use min hamming distance to determine the key size
// returning a slice of int with size of limit, the sizes that has lowest humming distance
func findKeySizes(input string, limit int) []int {
	var candidates []keyScore

	// trying out key size from 2 to 40
	for i := 2; i <= 40; i++ {

		// for each key size, calculate avg of humming distance of first 4 chunks
		sumDistance := float32(0)
		var counter int
		var str1 string
		var str2 string
		for j := 0; j < 4; j++ {
			rightEnd := i * (j + 2)
			if rightEnd >= len(input) {
				break
			}
			str1 = input[i*j : i*(j+1)]
			str2 = input[i*(j+1) : i*(j+2)]
			distance := hammingDistance(str1, str2)
			sumDistance += float32(distance) / float32(i)
			counter++
		}

		normalizedSum := sumDistance
		avgDistance := normalizedSum / float32(counter)
		candidates = append(candidates, keyScore{size: i, score: avgDistance})

		fmt.Printf("avg humming distance is %f when key size is %d\n", avgDistance, i)
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score < candidates[j].score
	})

	lowestHummingDistanceSizes := make([]int, limit)
	for i := 0; i < limit; i++ {
		lowestHummingDistanceSizes[i] = candidates[i].size
	}
	return lowestHummingDistanceSizes
}

func hammingDistance(str1 string, str2 string) int {
	if len(str1) < len(str2) {
		temp := str1
		str1 = str2
		str2 = temp
	}

	var sum int
	for i, c1 := range str1 {
		if i >= len(str2) {
			sum += countBits(byte(c1))
		} else {
			c2 := str2[i]
			sum += countBits(byte(c1) ^ c2)
		}
	}
	return sum
}

func countBits(b byte) int {
	var count int
	for i := 0; i < 8; i++ {
		if b&(1<<i) != 0 {
			count++
		}
	}
	return count
}

// transform the input bytes into chunks of size keySize
// discard incomplete chunk
func chunkBytes(input []byte, keySize int) [][]byte {
	var chunks [][]byte
	for i := 0; i+keySize <= len(input); i += keySize {
		chunk := input[i : i+keySize]
		chunks = append(chunks, chunk)
	}
	return chunks
}

func transpose(blocks [][]byte, keySize int) [][]byte {
	transposed := make([][]byte, keySize)

	for i := 0; i < keySize; i++ {
		idxSlice := make([]byte, len(blocks))
		for j := 0; j < len(blocks); j++ {
			idxSlice[j] = blocks[j][i]
		}
		transposed[i] = idxSlice
	}
	return transposed
}
