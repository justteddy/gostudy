package main

import (
	"bytes"
	"fmt"
	"io"
)

// IOLimitReader limited reader
type IOLimitReader struct {
	r     io.Reader
	limit int64
}

func (l *IOLimitReader) Read(p []byte) (n int, err error) {
	if l.limit <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.limit {
		p = p[0:l.limit]
	}
	n, err = l.r.Read(p)
	l.limit -= int64(n)
	return
}

func (l *IOLimitReader) Size() int64 {
	return l.limit
}

// LimitReader for limited byte count reading
func LimitReader(r io.Reader, n int64) io.Reader {
	lr := IOLimitReader{r, n}

	return &lr
}

func main() {
	reader := LimitReader(bytes.NewReader([]byte("some_string")), 5)
	buffer := make([]byte, 20)
	n, err := reader.Read(buffer)
	fmt.Println(n, err, string(buffer), buffer)
}
