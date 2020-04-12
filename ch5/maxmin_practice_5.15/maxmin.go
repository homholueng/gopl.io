package main

import (
	"fmt"
)

func compare(repl func(int, int) bool, vals ...int) int {
	if len(vals) == 0 {
		panic("vals can not be empty")
	}
	result := vals[0]
	for _, val := range vals[1:] {
		if repl(result, val) {
			result = val
		}
	}
	return result
}

func max(vals ...int) int {
	return compare(func(v1, v2 int) bool {
		return v1 < v2
	}, vals...)
}

func min(vals ...int) int {
	return compare(func(v1, v2 int) bool {
		return v1 > v2
	}, vals...)
}

func max1(val int, vals ...int) int {
	return compare(func(v1, v2 int) bool {
		return v1 < v2
	}, append(vals, val)...)
}

func min1(val int, vals ...int) int {
	return compare(func(v1, v2 int) bool {
		return v1 > v2
	}, append(vals, val)...)
}

func main() {
	fmt.Println(max(1, 2, 3, 4, 5))
	fmt.Println(min(1, 2, 3, 4, 5))
	fmt.Println(max1(1, 2, 3, 4, 5))
	fmt.Println(min1(1, 2, 3, 4, 5))
}
