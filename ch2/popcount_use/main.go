package main

import (
	"fmt"

	"github.com/homholueng/gopl.io/ch2/popcount"
)

func main() {
	fmt.Printf("%v %v\n", popcount.PopCount(1), popcount.PopCountMove(1))
	fmt.Printf("%v %v\n", popcount.PopCount(2), popcount.PopCountMove(2))
	fmt.Printf("%v %v\n", popcount.PopCount(31), popcount.PopCountMove(31))

}
