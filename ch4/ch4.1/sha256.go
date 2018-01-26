package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	var args = os.Args[1:]
	if len(args) >= 2 {
		n := diffBits(args[0], args[1])
		fmt.Println(n)
	}
}

func diffBits(str1, str2 string) int {
	c1 := sha256.Sum256([]byte(str1))
	c2 := sha256.Sum256([]byte(str2))

	n := 0
	for i := range c1 {
		for b := c1[i] ^ c2[i]; b != 0; b &= b - 1 {
			n++
		}
	}

	return n
}
