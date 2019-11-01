package main

import "fmt"

func main() {
	var sliceA = []string{"a", "b"}
	var arrayA = [3]string{"1", "2", "3"}
	for i := 0; i < len(sliceA); i++ {
		fmt.Println(sliceA[i])
	}
	for k, v := range arrayA {
		fmt.Println(k, v)
	}
}
