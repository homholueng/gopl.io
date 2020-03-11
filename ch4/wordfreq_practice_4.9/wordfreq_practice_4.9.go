package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	for {
		more := scanner.Scan()
		if !more {
			err := scanner.Err()
			if err != nil {
				fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
				os.Exit(1)
			}
			break
		}
		text := scanner.Text()
		words[text]++
	}
	fmt.Printf("word\tcount\n")
	for word, count := range words {
		fmt.Printf("%v\t%d\n", word, count)
	}
}
