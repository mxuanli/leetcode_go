package main

func shipWithinDays(weights []int, D int) int {
	start, end := 0, 0
	for _, w := range weights {
		if w > start {
			start = w
		}
		end += w
	}
	for start < end {
		mid := start + (end-start)/2
		days := 1
		tmp := 0
		for _, w := range weights {
			if tmp+w > mid {
				days += 1
				tmp = 0
			}
			tmp += w
		}
		if days <= D {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return end
}
