package main

import (
	"strings"
)

func main() {
	result := expand("$foobar can be really $foo! ", plusDollar)
	println(result)
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func plusDollar(s string) string {
	return "new"
}
