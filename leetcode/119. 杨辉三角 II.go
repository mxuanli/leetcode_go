package main

import "fmt"

func getRow(rowIndex int) []int {
	tmp := make([]int, rowIndex+1)
	tmp[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			tmp[j] = tmp[j] + tmp[j-1]
		}
	}
	return tmp
}

func main() {
	rowIndex := 3
	r := getRow(rowIndex)
	fmt.Println(r)
}
