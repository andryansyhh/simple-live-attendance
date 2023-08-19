package utils

import (
	"math/rand"
)

func CalculateSimilarity() (float64, error) {
	// for now  use random int to generate percentage of similarity,
	// in the future can use the third-party or library
	randomSimilarity := float64(rand.Intn(100))
	return randomSimilarity, nil
}
