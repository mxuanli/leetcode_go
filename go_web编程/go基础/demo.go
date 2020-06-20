package main

import "fmt"

func main() {
	a := []string{"a", "b", "C"}
	for i, j := range a {
		fmt.Println(i, j)
	}
}
