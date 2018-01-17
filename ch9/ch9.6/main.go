package main

import (
	"fmt"
	"runtime"
)

func main() {
	//get num cpu
	fmt.Println(runtime.NumCPU())

	runtime.GOMAXPROCS(6)

	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
