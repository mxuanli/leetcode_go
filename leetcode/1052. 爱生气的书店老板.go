package main

import "fmt"

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)
	start, end := 0, 0
	max0, tmp, r := 0, 0, 0
	for ; end < n; end++ {
		if grumpy[end] == 0 {
			r += customers[end]
		} else {
			tmp += customers[end]
		}
		if end-start+1 == X {
			if max0 < tmp {
				max0 = tmp
			}
			if grumpy[start] == 1 {
				tmp -= customers[start]
			}
			start += 1
		}
	}
	return r + max0
}

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	X := 3
	r := maxSatisfied(customers, grumpy, X)
	fmt.Println(r)
}
