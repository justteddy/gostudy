package main

import (
	"fmt"

	"gopl.io/ch2/tempconv"
)

func main() {
	fmt.Printf("3 celsius is - %v\n", tempconv.CtoK(3))
	fmt.Printf("3 kelvin is - %v\n", tempconv.KtoC(3))
	fmt.Printf("3 fahrenheit is - %v\n", tempconv.FtoK(3))
	fmt.Printf("3 kelvin is - %v\n", tempconv.KtoF(3))
}
