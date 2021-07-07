package main

func countPairs(deliciousness []int) int {
	maxD := deliciousness[0]
	for _, val := range deliciousness[1:] {
		if val > maxD {
			maxD = val
		}
	}
	mapD := make(map[int]int)
	maxD2 := maxD * 2
	r := 0
	for _, val := range deliciousness {
		for i := 1; i <= maxD2; i <<= 1 {
			r += mapD[i-val]
		}
		mapD[val]++
	}
	return r % (1e9 + 7)
}
