package main

func majorityElement(nums []int) int {
	major := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, num := range nums {
		if num == major {
			count++
		}
	}
	if count > len(nums)/2 {
		return major
	}
	return -1
}
