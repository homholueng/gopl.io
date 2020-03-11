package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4}
	a2 := []int{1, 2, 3, 4, 5}
	rotate(&a1)
	rotate(&a2)
	fmt.Println(a1)
	fmt.Println(a2)
}

func rotate(s *[]int) {
	l := len(*s)
	half := l / 2
	for i := 0; i < half; i++ {
		j := i + half + l%2
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
