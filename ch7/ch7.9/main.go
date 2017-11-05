package main

import "sort"
import "fmt"
import "strings"

func main() {
	x := []int{1, 2, 3, 2, 1}
	y := strings.Split("abcddcba", "")
	v := strings.Split("abcdbdcba", "")
	fmt.Println(IsPalindrome(sort.IntSlice(x)))
	fmt.Println(IsPalindrome(sort.StringSlice(y)))
	fmt.Println(IsPalindrome(sort.StringSlice(v)))
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
		} else {
			return false
		}
	}
	return true
}
