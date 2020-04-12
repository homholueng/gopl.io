package main

import (
	"fmt"
	"runtime"
)

func panicReturn() (num int) {
	defer func() {
		p := recover()
		fmt.Println(p)
		num = -1
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		fmt.Println(string(buf[:n]))
	}()
	panic("panic attack!")
}

func main() {
	fmt.Println(panicReturn())
}
