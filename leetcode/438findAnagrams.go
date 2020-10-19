package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var result []int
	if len(s) == 0 {
		return result
	}
	var frep [256]int
	for i := 0; i < len(p); i++ {
		frep[p[i]-'a']++
	}
	left, right, count := 0, 0, len(p)
	for right < len(s) {
		if frep[s[right]-'a'] >= 1 {
			count--
		}
		frep[s[right]-'a']--
		right++
		if count == 0 {
			result = append(result, left)
		}
		if right-left == len(p) {
			if frep[s[left]-'a'] >= 0 {
				count++
			}
			frep[s[left]-'a']++
			left++
		}
	}
	return result
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	result := findAnagrams(s, p)
	fmt.Println(result)
}
