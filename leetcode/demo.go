package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	var freq [256]int
	for i := 0; i < len(t); i++ {
		freq[t[i]-'a']++
	}
	left, right, lastLeft, LastRight, count := 0, 0, -1, len(s), len(t)
	for ; right < len(s); right++ {
		if freq[s[right]-'a'] >= 1 {
			count--
		}
		freq[s[right]-'a']--
		for count == 0 {
			if right-left < LastRight-lastLeft {
				lastLeft = left
				LastRight = right
			}
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
	}
	if lastLeft != -1 {
		return s[lastLeft : LastRight+1]
	}
	return ""
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	r := minWindow(s, t)
	fmt.Println(r)
}
