package main

import (
	"fmt"
	"sort"
)

func isMonotonic(A []int) bool {
	return sort.IntsAreSorted(A) || sort.IsSorted(sort.Reverse(sort.IntSlice(A)))
}

func main() {
	A := []int{1, 2, 2, 3}
	r := isMonotonic(A)
	fmt.Println(r)
}
