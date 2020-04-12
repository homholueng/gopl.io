package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter will count numbers of words and lines in bytes
type WordCounter struct {
	words int
	lines int
}

func (wc *WordCounter) Write(p []byte) (int, error) {
	wordScan := bufio.NewScanner(bytes.NewReader(p))
	lineScan := bufio.NewScanner(bytes.NewReader(p))

	wordScan.Split(bufio.ScanWords)
	for wordScan.Scan() {
		wc.words++
	}

	lineScan.Split(bufio.ScanLines)
	for lineScan.Scan() {
		wc.lines++
	}

	return len(p), nil
}

// Words return number of words wc scan
func (wc *WordCounter) Words() int {
	return wc.words
}

// Lines return number of lines wc scan
func (wc *WordCounter) Lines() int {
	return wc.lines
}

func main() {
	wc := WordCounter{}
	fmt.Fprintf(&wc, "a bc e f\n a \n a b c")
	fmt.Println(wc.Words(), wc.Lines())
}
