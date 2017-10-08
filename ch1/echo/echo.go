package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], ", "))
	fmt.Println("Ex 1 time = ", time.Since(start).Nanoseconds())
}

func ex2() {
	start := time.Now()
	for key, value := range os.Args[1:] {
		fmt.Println(key, value)
	}
	fmt.Println("Ex 2 time = ", time.Since(start).Nanoseconds())
}

func ex3() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("Ex 3 time = ", time.Since(start).Nanoseconds())
}
