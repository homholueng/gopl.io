package main

import "fmt"

func main() {
	fmt.Println(unique([]string{"a", "a", "b", "c", "c", "d"}))
}

func unique(strs []string) []string {
	w := 0
	for _, s := range strs {
		if strs[w] == s {
			continue
		}
		w++
		strs[w] = s
	}
	return strs[:w+1]
}
