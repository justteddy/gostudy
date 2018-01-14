package main

import "fmt"

type wd struct {
	amount int
	resch  chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan wd) // withdraw balance

func Deposit(amount int) {
	deposits <- amount
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	wdoper := wd{amount, ch}
	withdraws <- wdoper
	return <-ch
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case wdoper := <-withdraws:
			if balance < wdoper.amount {
				wdoper.resch <- false
			} else {
				balance = balance - wdoper.amount
				wdoper.resch <- true
			}
		case balances <- balance:
		}
	}
}

func main() {
	go teller() // start the monitor goroutine

	done := make(chan struct{})

	go func() {
		Deposit(400)
		if ok := Withdraw(250); ok {
			fmt.Println("Withdraw success")
		} else {
			fmt.Println("Withdraw failed")
		}
		done <- struct{}{}
	}()

	go func() {
		Deposit(100)
		if ok := Withdraw(50); ok {
			fmt.Println("Withdraw success")
		} else {
			fmt.Println("Withdraw failed")
		}
		done <- struct{}{}
	}()

	go func() {
		Deposit(300)
		if ok := Withdraw(300); ok {
			fmt.Println("Withdraw success")
		} else {
			fmt.Println("Withdraw failed")
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	<-done

	// always 200
	fmt.Println("Result balance =", Balance())
}
