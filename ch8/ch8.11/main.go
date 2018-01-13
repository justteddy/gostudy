package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type file struct {
	filename string
	url      string
	size     int64
}

var done = make(chan struct{})

func main() {

	responses := make(chan file, 3)
	for _, url := range os.Args[1:] {
		go func(url string) {
			fmt.Println("Started goroutine with")
			local, n, err := fetch(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
				return
			}
			responses <- file{local, url, n}
		}(url)
	}

	for {
		select {
		case <-done:
			f := <-responses
			fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", f.url, f.filename, f.size)
			return
		}
	}
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func fetch(url string) (filename string, n int64, err error) {
	if cancelled() {
		return
	}
	resp, err := http.Get(url)
	fmt.Printf("request sended to %s\n", url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	//stop all goroutines
	close(done)

	local := strings.TrimPrefix(strings.Replace(url, "/", "", -1), "https:") + ".html"
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
