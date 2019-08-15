package main

import "fmt"

func main() {
	a := 10
	b := 0
	testIf(a, b)
}

func testIf(a int, b int) {
	if a < b {
		fmt.Println("a小于b")
	} else if a > b {
		fmt.Println("a大于b")
	}
}
