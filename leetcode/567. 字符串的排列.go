package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	start, end, count, n1, n2 := 0, 0, len(s1), len(s1), len(s2)
	strMap := make(map[uint8]int)
	for i := 0; i < n1; i++ {
		tmp := s1[i]
		strMap[tmp]++
	}
	for ; end < n2; end++ {
		tmpEnd := s2[end]
		if strMap[tmpEnd] > 0 {
			count--
		}
		strMap[tmpEnd]--
		if count == 0 {
			return true
		}
		if end-start+1 == n1 {
			tmpStart := s2[start]
			if strMap[tmpStart] >= 0 {
				count++
			}
			strMap[tmpStart]++
			start++
		}
	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	r := checkInclusion(s1, s2)
	fmt.Println(r)
}
