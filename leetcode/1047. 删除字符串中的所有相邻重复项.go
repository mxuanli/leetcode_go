package main

import (
	"fmt"
)

func removeDuplicates(S string) string {
	var stack []int32
	for _, s := range S {
		if len(stack) != 0 && stack[len(stack)-1] == s {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s)
		}
	}
	return string(stack)
}

func main() {
	S := "abbaca"
	r := removeDuplicates(S)
	fmt.Println(r)
}
