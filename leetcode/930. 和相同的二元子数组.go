package main

func numSubarraysWithSum(nums []int, goal int) int {
	numsMap := make(map[int]int)
	sum := 0
	r := 0
	for _, num := range nums {
		numsMap[sum]++
		sum += num
		r += numsMap[sum-goal]
	}
	return r
}
