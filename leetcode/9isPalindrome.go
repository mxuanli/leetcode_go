package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	lenX := len(strX)
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	for i := 0; i < lenX/2; i++ {
		if strX[i] != strX[lenX-i-1] {
			return false
		}
	}
	return true
}

func main() {
	x := -121
	r := isPalindrome(x)
	fmt.Println(r)
}
