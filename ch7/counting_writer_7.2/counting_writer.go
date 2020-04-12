package main

import (
	"fmt"
	"io"
)

type dummyWriter struct {
}

func (dw dummyWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type countingWriter struct {
	count  int64
	writer io.Writer
}

func (cw *countingWriter) Write(p []byte) (n int, err error) {
	n, err = cw.writer.Write(p)
	cw.count += int64(n)
	return
}

// CountingWriter return a writer and a count
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := countingWriter{writer: w}
	return &cw, &cw.count
}

func main() {
	w := dummyWriter{}
	cw, count := CountingWriter(w)
	cw.Write([]byte{1, 2, 3})
	cw.Write([]byte{1, 2, 3})
	cw.Write([]byte{1, 2, 3})
	fmt.Println(*count)
}
