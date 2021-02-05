package main

import "fmt"

func sum(Slice []int) int {
	r := 0
	for _, s := range Slice {
		r = r + s
	}
	return r
}

func findMaxAverage2(nums []int, k int) float64 {
	maxSum, sumNums := sum(nums[0:k]), sum(nums[0:k])
	n := len(nums)
	for i := k; i < n; i++ {
		start := nums[i-k]
		end := nums[i]
		sumNums = sumNums - start + end
		if sumNums > maxSum {
			maxSum = sumNums
		}
	}
	return float64(maxSum) / float64(k)
}

func findMaxAverage(nums []int, k int) float64 {
	maxSum, sumNums, n := sum(nums[0:k]), sum(nums[0:k]), len(nums)
	for i := k; i < n; i++ {
		sumNums = sumNums - nums[i-k] + nums[i]
		if sumNums > maxSum {
			maxSum = sumNums
		}
	}
	return float64(maxSum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	r := findMaxAverage(nums, k)
	fmt.Println(r)
}
