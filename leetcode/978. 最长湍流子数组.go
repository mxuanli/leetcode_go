package main

import "fmt"

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n == 1 {
		return n
	}
	if n == 2 {
		if arr[0] == arr[1] {
			return 1
		}
		return n
	}
	count, r := 0, 0
	if arr[0] == arr[1] {
		count = 1
	} else {
		count = 2
	}
	for i := 2; i < n; i++ {
		if (arr[i-2] > arr[i-1] && arr[i-1] < arr[i]) || (arr[i-2] < arr[i-1] && arr[i-1] > arr[i]) {
			count++
		} else {
			if arr[i-1] == arr[i] {
				count = 1
			} else {
				count = 2
			}
		}
		if count > r {
			r = count
		}
	}
	return r
}

func main() {
	arr := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	r := maxTurbulenceSize(arr)
	fmt.Println(r)
}
