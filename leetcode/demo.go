package main

import "fmt"

func main() {
	nums := 123
	var sliceN []int
	for nums != 0 {
		tmp := nums % 10
		nums = nums / 10
		sliceN = append(sliceN, tmp)
	}
	fmt.Println(sliceN)
}
