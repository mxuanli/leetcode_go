package main

import "fmt"

func countBits(num int) []int {
	var r []int
	for i := 0; i <= num; i++ {
		count := 0
		j := i
		for j != 0 {
			j = j & (j - 1)
			count++
		}
		r = append(r, count)
	}
	return r
}

func main() {
	num := 3
	r := countBits(num)
	fmt.Println(r)
}
