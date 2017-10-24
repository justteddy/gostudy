package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	use(&x)

	fmt.Println(x)
}

func use(x *[]int) {
	slice := *x
	*x = append(slice[:1], slice[2:]...)
	// x = append(x[:1], x[2:]...)
}
