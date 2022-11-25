package main

import (
	"testing"
)

func Benchmark_Clear(b *testing.B) {
	b.Run("clear", func(b *testing.B) {
		// data intitialization here
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// clear all invalid book orders here
		}
	})
}
