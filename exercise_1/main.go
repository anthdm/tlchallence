package main

import (
	"math"
)

const (
	MaxFactor = 1e8
	LastPrice = 16551.2
)

func main() {
	// var (
	// 	asksSrc = "../data/asks.gob"
	// 	bidSrc  = "../data/bids2.gob"
	// 	book    = orderbook.NewFromFile(asksSrc, bidSrc)
	// )

	// The two sides of the book are OUT of sync
	// (eg not receiving updates due to a network error).
	// Assuming that the latest order took place at "LastPrice"
	// Find a fast way to clear the invalid orders on both sides of the book.

	// Do not use the original maps, convert them to a different data structure
	// that is more suitable to solve the problem.

	// Write a benchmark and a test to validate your solution.

	// Note: You are allowed to use any data type (ex: int64*MaxFactor instead of float64), data structure or existing libraries.
	// Note: Initial data structure creation is not part of the benchmark. Only the clearing operation is.
}

// Round rounds a float64 to the provided precision factor.
func Round(n float64) float64 {
	return math.Round(n*MaxFactor) / MaxFactor
}
