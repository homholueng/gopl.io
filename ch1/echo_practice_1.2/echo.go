package main

import (
	"fmt"
	"os"
)

func main() {
	for index, val := range os.Args {
		fmt.Printf("%d %v\n", index, val)
	}
}
