package main

import "fmt"

func longestOnes(A []int, K int) int {
	start, end, count0, count1, r, n := 0, 0, 0, 0, 0, len(A)
	for ; end < n; end++ {
		if A[end] == 1 {
			count1++
		}
		if A[end] == 0 {
			count0++
		}
		for count0 > K {
			if A[start] == 1 {
				count1--
			}
			if A[start] == 0 {
				count0--
			}
			start++
		}
		if r < end-start+1 {
			r = end - start + 1
		}
	}
	return r
}

func main() {
	A := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	K := 2
	r := longestOnes(A, K)
	fmt.Println(r)
}
