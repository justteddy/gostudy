package main

import (
	"fmt"
	"os"
)

func main() {
	for key, value := range os.Args[1:] {
		fmt.Println(key, value)
	}
}
