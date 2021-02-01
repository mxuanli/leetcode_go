package main

import "fmt"

func main() {
	A := []int{1, 1}
	for i, j := range A {
		fmt.Println(i, j)
	}
}
