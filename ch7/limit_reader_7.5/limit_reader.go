package main

import (
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	limit  int64
	reader io.Reader
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	_p := make([]byte, lr.limit)
	n, err = lr.reader.Read(_p)
	if err != nil {
		return
	}
	copy(p, _p)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{limit: n, reader: r}
}

func main() {
	reader := strings.NewReader("Haha")
	lr := LimitReader(reader, 1)
	fmt.Println(reader.Len())
	b := make([]byte, 4)
	lr.Read(b)
	fmt.Println(b)
}
