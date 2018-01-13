package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

func handleConn(conn net.Conn) {
	who := conn.RemoteAddr().String()
	ch := make(chan string)
	me := client{ch, who}
	go clientWriter(conn, ch)

	me.channel <- "You are " + me.name
	messages <- me.name + " has arrived"
	entering <- me

	input := bufio.NewScanner(conn)
	for input.Scan() {
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
