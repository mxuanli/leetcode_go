package main

func totalHammingDistance(nums []int) int {
	n := len(nums)
	r := 0
	for i := 0; i < 30; i++ {
		c := 0
		for _, num := range nums {
			c += (num >> i) & 1
		}
		r += c * (n - c)
	}
	return r
}
