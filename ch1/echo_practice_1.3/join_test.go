package bench

import (
	"strings"
	"testing"
)

func forJoin(args []string) string {
	var sep, s string
	for _, val := range args {
		s += sep + val
		sep = " "
	}
	return s
}

func stringJoin(args []string) string {
	return strings.Join(args, " ")
}

func BenchmarkForJoin(b *testing.B) {
	args := []string{"1", "2", "3", "4"}
	for i := 0; i < b.N; i++ {
		forJoin(args)
	}
}

func BenchmarkStringJoin(b *testing.B) {
	args := []string{"1", "2", "3", "4"}
	for i := 0; i < b.N; i++ {
		stringJoin(args)
	}
}
