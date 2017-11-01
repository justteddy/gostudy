package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter counts words
type WordCounter int

// LineCounter counts lines
type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*w++
	}

	return int(*w), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*l++
	}

	return int(*l), nil
}

func main() {
	var w WordCounter
	var l LineCounter

	w.Write([]byte("hello mate, it's only 7 words here!"))
	fmt.Printf("%d words here\n", w)

	l.Write([]byte("hello mate\n it's only \n3 lines here!"))
	fmt.Printf("%d lines here\n", l)
}
