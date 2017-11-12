package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack [][]string

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			nameAndAttributes := []string{}
			nameAndAttributes = append(nameAndAttributes, tok.Name.Local)
			for _, val := range tok.Attr {
				nameAndAttributes = append(nameAndAttributes, strings.Join([]string{val.Name.Local, val.Value}, "="))
			}
			stack = append(stack, nameAndAttributes) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				for _, nameAndAttributes := range stack {
					fmt.Printf("%s ", strings.Join(nameAndAttributes, "|"))
				}
				fmt.Printf(": %s\n", tok)
			}
		}
	}
}

func containsAll(x [][]string, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		for _, element := range x[0] {
			if element == y[0] {
				y = y[1:]
			}
		}
		x = x[1:]
	}
	return false
}
