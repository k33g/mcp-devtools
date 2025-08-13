package rag

import (
	"math"
	"sort"
)

// getTopNVectorRecords returns the top N vector records based on their cosine similarity.
func getTopNVectorRecords(records []VectorRecord, max int) []VectorRecord {
	// Sort the records slice in descending order based on CosineDistance
	sort.Slice(records, func(i, j int) bool {
		return records[i].CosineSimilarity > records[j].CosineSimilarity
	})

	// Return the first max records or all if less than three
	if len(records) < max {
		return records
	}
	return records[:max]
}

// --- Cosine similarity ---

// dotProduct calculates the dot product of two vectors
// It assumes that both vectors are of the same length.
func dotProduct(v1 []float64, v2 []float64) float64 {
	// Calculate the dot product of two vectors
	sum := 0.0
	for i := range v1 {
		sum += v1[i] * v2[i]
	}
	return sum
}

// cosineSimilarity calculates the cosine similarity between two vectors
func cosineSimilarity(v1, v2 []float64) float64 {
	// Calculate the cosine distance between two vectors
	product := dotProduct(v1, v2)

	norm1 := math.Sqrt(dotProduct(v1, v1))
	norm2 := math.Sqrt(dotProduct(v2, v2))
	if norm1 <= 0.0 || norm2 <= 0.0 {
		// Handle potential division by zero
		return 0.0
	}
	return product / (norm1 * norm2)
}
