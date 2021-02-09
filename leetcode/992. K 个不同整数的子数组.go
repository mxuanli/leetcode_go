package main

import "fmt"

func foo(A []int, K int) int {
	// 由最多K个不同整数组成的子数组的个数
	n := len(A)
	start, end, nums, r := 0, 0, 0, 0
	numsMap := make(map[int]int)
	for ; end < n; end++ {
		num := A[end]
		if numsMap[num] == 0 {
			nums++
		}
		numsMap[num]++
		for nums > K {
			num := A[start]
			numsMap[num]--
			if numsMap[num] == 0 {
				nums--
			}
			start++
		}
		r += end - start + 1
	}
	return r
}

func subarraysWithKDistinct(A []int, K int) int {
	num1 := foo(A, K)
	num2 := foo(A, K-1)
	return num1 - num2
}

func main() {
	A := []int{1, 2, 1, 2, 3}
	K := 2
	r := subarraysWithKDistinct(A, K)
	fmt.Println(r)
}
