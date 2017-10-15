package main

import "fmt"

func main() {
	slice := []string{"some", "string", "here"}
	fmt.Println(join(" KEK!", slice...))
}

func join(sep string, vals ...string) string {
	var result string

	length := len(vals)
	for i, val := range vals {
		result = result + val
		if i != length-1 {
			result = result + sep
		}
	}

	return result
}
