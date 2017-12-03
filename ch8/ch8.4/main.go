package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go func(ch chan int) {
		for i := 1; i <= 4; i++ {
			ch <- i
			fmt.Printf("%d Element added\n", i)
		}
		close(ch)
	}(ch)

	time.Sleep(300 * time.Millisecond)
	fmt.Println(len(ch))
	<-ch
	fmt.Println(len(ch))
	<-ch
	fmt.Println(len(ch))

	// time.Sleep(5000 * time.Millisecond)
	// fmt.Println("Can't load 4-th element, until read from channel")

	// for v := range ch {
	// 	fmt.Println("Done", v)
	// }
}
