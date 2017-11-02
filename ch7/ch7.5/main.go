package main

import "flag"
import "fmt"

type CustomFlag struct {
	s string
}

func (c *CustomFlag) String() string {
	return c.s
}

func (c *CustomFlag) Set(text string) error {

	switch text {
	case "ny":
		c.s = "New-York"
	case "nsk":
		c.s = "Novosibirsk"
	default:
		c.s = "Moscow"
	}

	return nil
}

func main() {
	cflag := CustomFlag{}
	flag.Var(&cflag, "city", "Enter some city name")
	flag.Parse()

	fmt.Printf("%s", cflag)
}
