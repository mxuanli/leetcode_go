package main

import "fmt"

func maxValue(tmp map[byte]int) int {
	r := 0
	for _, v := range tmp {
		if r < v {
			r = v
		}
	}
	return r
}

func characterReplacement(s string, k int) int {
	start, end, n := 0, 0, len(s)
	tmp := make(map[byte]int)
	r := 0
	for end = 0; end < n; end++ {
		tmp[s[end]]++
		v := maxValue(tmp)
		if (end - start + 1) > (v + k) {
			tmp[s[start]]--
			start++
		} else {
			if r < (end - start + 1) {
				r = end - start + 1
			}
		}
	}
	return r
}

func main() {
	s := "ABAB"
	k := 2
	r := characterReplacement(s, k)
	fmt.Println(r)
}
