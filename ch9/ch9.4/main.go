package main

import (
	"fmt"
	"time"
)

func f1() {
	ch := make(chan struct{})
	begin := time.Now()
	var count int64 = 0

	defer func(begin time.Time) {
		end := time.Now()
		fmt.Printf("\nDone after %d seconds\n", end.Sub(begin)/1000000)
		if p := recover(); p != nil {
			fmt.Print("Out of memory. Sending data to all goroutines...\n")
		}
	}(begin)

	fmt.Println("Number of goroutines:")
	for {
		count++
		fmt.Printf("\r%d", count)
		go func() { <-ch }()
	}
}

func f2() {
	chan1 := make(chan struct{})
	chan2 := make(chan struct{})
	chan3 := make(chan struct{})

	go func() {
		for i := 0; i < 100; i++ {
			chan1 <- struct{}{}
		}
		close(chan1)
	}()

	go func() {
		for range chan1 {
			chan2 <- struct{}{}
		}
		close(chan2)
	}()

	go func() {
		for range chan2 {
			chan3 <- struct{}{}
		}
		close(chan3)
	}()

	for range chan3 {
		fmt.Println("value")
	}
}
