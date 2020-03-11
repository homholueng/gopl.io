package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4}
	a2 := []int{1, 2, 3, 4, 5}
	reverse(&a1)
	reverse(&a2)
	fmt.Println(a1)
	fmt.Println(a2)
}

func reverse(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
