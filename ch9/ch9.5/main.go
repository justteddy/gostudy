package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(abort)
	}()

	go func() {
		for {
			start := time.Now()
			fmt.Println(<-ch2)
			ch1 <- "pong"
			end := time.Now()
			fmt.Println("Time since start - ", end.Sub(start))
		}
	}()

	go func() {
		for {
			ch2 <- "ping"
			fmt.Println(<-ch1)
		}
	}()

	<-abort
}
