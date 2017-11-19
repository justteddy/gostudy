package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	timezones := map[string]string{
		"localhost:8010": "europe",
		"localhost:8020": "asia",
		"localhost:8030": "america",
	}
	hosts := []string{
		"localhost:8010",
		"localhost:8020",
		"localhost:8030",
	}

	conns := make([]net.Conn, len(timezones))
	for i, host := range hosts {
		conn, err := net.Dial("tcp", host)
		if err != nil {
			log.Fatal(err)
		}
		conns[i] = conn
		defer conn.Close()
		fmt.Print(timezones[host])
		fmt.Print("\t\t\t")
	}
	fmt.Println()

	for {
		for i, conn := range conns {
			b := make([]byte, 1024)
			n, err := conn.Read(b)
			if err == io.EOF {
				os.Exit(0)
			}
			if err == nil && n > 0 {
				for j := 0; j < i; j++ {
					os.Stdout.WriteString("\t\t\t")
				}
				os.Stdout.Write(b)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
