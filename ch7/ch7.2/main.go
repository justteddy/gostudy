package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// CustomWriter obj
type CustomWriter struct {
	length int64
	w      io.Writer
}

func (c *CustomWriter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	if err != nil {
		log.Fatal("write error")
	}

	c.length = int64(n)
	return
}

// CountingWriter is a wrapper for writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CustomWriter{0, w}
	return &cw, &cw.length
}

func main() {
	writer, length := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "Hello world, bitch!\n")
	fmt.Println(*length)
}
