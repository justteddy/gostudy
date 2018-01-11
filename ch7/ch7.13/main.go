package main

import (
	"fmt"
)

type test struct{}

func (t test) testMethod() {
	fmt.Println(123)
}

type ti interface {
	testMethod()
}

func main() {
	// var w io.Writer
	// w = os.Stdout
	// rw, ok := w.(io.Writer)

	// var w interface{}
	// w = 3
	// rw, ok := w.(int)

	// var w io.Writer
	// w = os.Stdout
	// rw := w.(io.ReadWriter)

	// rw.Read([]byte("asdas"))

	// fmt.Printf("type w - %T, type rw - %T\n", w, rw)
	var intr interface{} = test{}
	res, ok := intr.(ti)
	fmt.Printf("%T %v", res, ok)
}
