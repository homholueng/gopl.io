package main

import (
	"fmt"
	"strings"
)

func stringJoin(sep string, a ...string) string {
	return strings.Join(a, sep)
}

func main() {
	fmt.Println(stringJoin(",", "a", "b", "c", "d"))
	fmt.Println(stringJoin(",", "a", "b"))
}
