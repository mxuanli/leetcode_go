package main

import (
	"fmt"
)

func main() {
	var stringA string = "hello"
	fmt.Println(stringA)
	var newString = new(string)
	fmt.Println(newString)
	*newString = "world"
	fmt.Println(*newString)
	mapA := make(map[string]string)
	fmt.Println(mapA)
	sliceA := make([]int, 2, 3)
	fmt.Println(sliceA)
}
