package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {

	fmt.Println(strings.Join(strings.Fields(string("x\t\t\t世\t\t\ty \v\v\n  z")), " "))
	fmt.Println(strings.Join(strings.Fields(string("Räk  smö rg å   s")), " "))
	fmt.Println(string(uniqueUtf8Space([]byte("x\t\t\t世\t\t\ty \v\v\n  z"))))
	fmt.Println(string(uniqueUtf8Space([]byte("Räk  smö rg å   s"))))
}

func uniqueUtf8Space(bytes []byte) []byte {
	i := 0
	runes := string(bytes)
	spaceAhaed := false
	for _, r := range runes {
		if unicode.IsSpace(r) {
			if spaceAhaed {
				continue
			} else {
				spaceAhaed = true
				r = ' '
			}
		} else {
			spaceAhaed = false
		}

		if utf8.RuneLen(r) > 1 {
			buf := make([]byte, 4)
			len := utf8.EncodeRune(buf, r)
			copy(bytes[i:i+len], buf)
			i += len
		} else {
			bytes[i] = byte(r)
			i++
		}
	}
	return bytes[:i]
}
