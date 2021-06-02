package main

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	pre := 0
	remainderMap := make(map[int]int)
	remainderMap[0] = -1
	for i, num := range nums {
		pre = (pre + num) % k
		if preIndex, ok := remainderMap[pre]; ok {
			if i-preIndex >= 2 {
				return true
			}
		} else {
			remainderMap[pre] = i
		}
	}
	return false
}
