package main

import "math"

func judgeSquareSum(c int) bool {
	start := 0
	end := int(math.Sqrt(float64(c))) + 1
	for start <= end {
		sum := start*start + end*end
		if sum == c {
			return true
		} else if sum < c {
			start += 1
		} else {
			end -= 1
		}
	}
	return false
}
