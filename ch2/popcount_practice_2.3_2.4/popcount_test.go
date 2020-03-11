package bech

import (
	"testing"

	"github.com/homholueng/gopl.io/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(231)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(231)
	}
}

func BenchmarkPopCountMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountMove(231)
	}
}

func BenchmarkPopCountClearLowBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountClearLowBit(231)
	}
}
