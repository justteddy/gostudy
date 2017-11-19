package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn, addedhours int) {
	defer c.Close()
	for {
		timein := time.Now().Local().Add(time.Hour * time.Duration(addedhours))
		_, err := io.WriteString(c, timein.Format("15:04:05\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8010", "port to listen")
	flag.Parse()

	var addedhours int
	switch port {
	case "8010":
		addedhours = 2
	case "8020":
		addedhours = 4
	case "8030":
		addedhours = 6
	}

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, addedhours)
	}
}
