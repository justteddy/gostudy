package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	channel chan<- string
	name    string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[string]*client) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for _, cli := range clients {
				cli.channel <- msg
			}

		case cli := <-entering:
			notifyAboutClient(cli.channel, clients)
			clients[cli.name] = &cli

		case cli := <-leaving:
			delete(clients, cli.name)
			close(cli.channel)
		}
	}
}

func notifyAboutClient(ch chan<- string, clients map[string]*client) {
	ch <- "All active clients:"
	for name := range clients {
		ch <- name
	}
}

func timeoutCounter(conn net.Conn, chout <-chan bool) {
	ticker := time.NewTicker(time.Second)
	limit := 60 // 1 minute
	count := 0

	for {
		select {
		case <-ticker.C:
			count++
			if count == limit {
				msg := conn.RemoteAddr().String() + "away too long. Kicked out."
				messages <- msg
				fmt.Fprintln(conn, msg)
				ticker.Stop()
				conn.Close()
				return
			}
		case <-chout:
			count = 0
		}
	}
}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	fmt.Fprint(conn, "Enter your name:")

	var who string
	if input.Scan() {
		who = input.Text()
	}

	ch := make(chan string)
	me := client{ch, who}
	go clientWriter(conn, ch)

	chout := make(chan bool)
	go timeoutCounter(conn, chout)

	me.channel <- "You are " + me.name
	messages <- me.name + " has arrived"
	entering <- me

	for input.Scan() {
		chout <- true
		messages <- me.name + ": " + input.Text()
	}

	leaving <- me
	messages <- me.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
