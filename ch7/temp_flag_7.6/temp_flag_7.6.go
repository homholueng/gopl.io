package main

import (
	"flag"
	"fmt"

	"github.com/homholueng/gopl.io/ch7/tempconv"
)

var temp = tempconv.FahrenheitFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
