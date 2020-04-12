package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	i := 0
	j := s.Len() - 1
	for i < j {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		i++
		j--
	}
	return true
}

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	fmt.Println(IsPalindrome(StringSlice([]string{"a", "c", "b", "c", "a"})))
	fmt.Println(IsPalindrome(StringSlice([]string{"a", "c", "b", "b", "c", "a"})))
	fmt.Println(IsPalindrome(StringSlice([]string{"a", "b", "c", "d", "c", "b", "a"})))
	fmt.Println(IsPalindrome(StringSlice([]string{"a", "b", "c", "d", "e", "c", "b", "a"})))
}
