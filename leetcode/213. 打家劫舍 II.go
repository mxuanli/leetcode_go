package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	a, b := 0, 0
	for i := 1; i < len(nums); i++ {
		tmp := a
		a = b
		if b < tmp + nums[i] {
			b =  tmp + nums[i]
		}
	}
	r := b
	a, b = 0, 0
	for i := 0; i < len(nums) - 1; i++ {
		tmp := a
		a = b
		if b < tmp + nums[i] {
			b = tmp + nums[i]
		}
	}
	if r < b {
		r = b
	}
	return r
}
