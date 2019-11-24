package main

import "fmt"

func main() {
	FuncArray()
}

func FuncArray() {
	var arrayA = [3]string{"abc", "def", "ghi"}
	fmt.Println(arrayA)

	arrayB := [...]string{"jkl", "mno", "pqr", "stu"}
	for _, v := range arrayB {
		fmt.Print(v)
	}

	for i := range arrayB {
		fmt.Println(i)
	}
}
