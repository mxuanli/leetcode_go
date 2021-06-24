package main

func hammingWeight(num uint32) int {
	r := 0
	for num > 0 {
		num = num & (num - 1)
		r += 1
	}
	return r
}
