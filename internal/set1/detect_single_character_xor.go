package set1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// Detect from all the fixed length HEX strings in the file, which one is XOR'd, and decipher it
func Detect() (string, error) {
	file, err := os.Open("../../resources/challenge-data/4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	minLoss := math.MaxFloat64
	var candidateRes string

	for scanner.Scan() {
		line := scanner.Text()
		res, loss, _, err2 := Decipher(line, nil)
		if err2 != nil {
			return "", err2
		}
		if loss < minLoss {
			minLoss = loss
			candidateRes = res
			fmt.Printf("Candidate: %s, with loss %f\n", res, loss)
		}
	}
	return candidateRes, nil
}
