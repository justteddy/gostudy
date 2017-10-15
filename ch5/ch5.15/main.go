package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 20, 30, 18, 16, 15, 12, 100, 99, 103}
	res, ok := max(slice...)
	if !ok {
		fmt.Println("Error!")
	}

	fmt.Printf("Max num is - %d\n", res)

	res, ok = min(slice...)
	if !ok {
		fmt.Println("Error!")
	}

	fmt.Printf("Min num is - %d\n", res)
}

func max(values ...int) (result int, ok bool) {

	if len(values) == 0 {
		return
	}

	ok = true
	result = math.MinInt64
	for _, val := range values {
		if val > result {
			result = val
		}
	}
	return
}

func min(values ...int) (result int, ok bool) {
	if len(values) == 0 {
		return
	}

	ok = true
	result = math.MaxInt64
	for _, val := range values {
		if val < result {
			result = val
		}
	}
	return
}
