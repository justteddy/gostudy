package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var args = os.Args[2:]

	shaSize := flag.Int("size", 256, "sha size format")
	flag.Parse()

	for _, arg := range args {
		switch *shaSize {
		case 256:
			result := sha256.Sum256([]byte(arg))
			fmt.Printf("%x\n", result)
		case 384:
			result := sha512.Sum384([]byte(arg))
			fmt.Printf("%x\n", result)
		case 512:
			result := sha512.Sum512([]byte(arg))
			fmt.Printf("%x\n", result)
		}
	}
}
