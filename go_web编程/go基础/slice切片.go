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
	Slice()
}

func Slice() {
	var sliceC []string
	sliceC = append(sliceC, "hello") // 向slice里追加值
	fmt.Println(sliceC)
	var sliceD = []string{"abc"}
	fmt.Println(sliceD)
}
