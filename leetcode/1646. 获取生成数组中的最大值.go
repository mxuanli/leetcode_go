package main

import "fmt"

func getMaximumGenerated(n int) int {
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	for i := 0; i <= n+1; i++ {
		i2 := 2 * i
		if (2 == i2 || 2 < i2) && i2 == n || i2 < n {
			nums[i2] = nums[i]
		}
		if (2 == i2+1 || 2 < i2+1) && (i2+1 == n || i2+1 < n) {
			nums[i2+1] = nums[i] + nums[i+1]
		}
	}
	r := 0
	for _, num := range nums {
		if r < num {
			r = num
		}
	}
	fmt.Println(nums)
	return r
}

func main() {
	getMaximumGenerated(10)
}
