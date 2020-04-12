package main

import (
	"fmt"

	"github.com/homholueng/gopl.io/ch6/intset"
)

func main() {

	set1 := intset.IntSet{}
	set1.Add(1)
	set1.Add(64)
	set1.Add(123)
	fmt.Printf("len of set1(%v): %v\n", set1, set1.Len())
	set1.Remove(3)
	fmt.Printf("set1 after remove 3: %v\n", set1)
	set1.Remove(64)
	fmt.Printf("set1 after remove 64: %v\n", set1)
	set2 := set1.Copy()
	fmt.Printf("copy of set1: %v\n", set2)
	set1.Clear()
	fmt.Printf("set1 after clear: %v\n", set1)
	fmt.Printf("set2 after clear: %v\n", set2)

	set1.AddAll(1, 64, 128, 129, 200)
	fmt.Printf("set1 after addall: %v\n", set1)

	set1.Clear()
	set2.Clear()
	set1.AddAll(1, 64)
	set2.AddAll(2, 64, 65, 128, 256)
	fmt.Printf("set1(%v) intersct set2(%v): ", set1, set2)
	set1.IntersectWith(set2)
	fmt.Printf("%v\n", set1)

	set1.Clear()
	set2.Clear()
	set1.AddAll(1, 64, 67)
	set2.AddAll(2, 64, 65, 128, 256)
	fmt.Printf("set1(%v) difference set2(%v): ", set1, set2)
	set1.DifferenceWith(set2)
	fmt.Printf("%v\n", set1)

	set1.Clear()
	set2.Clear()
	set1.AddAll(1, 64, 67, 128)
	set2.AddAll(2, 64, 65, 128, 256)
	fmt.Printf("set1(%v) symmetric difference set2(%v): ", set1, set2)
	set1.SymmetricDifference(set2)
	fmt.Printf("%v\n", set1)

	set1.Clear()
	set1.AddAll(1, 64, 67, 128)
	fmt.Printf("set1(%v) elements: %v", set1, set1.Elems())
}
