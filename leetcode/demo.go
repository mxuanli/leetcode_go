package main

import "fmt"

func minWindow(s string, t string) string {
	result := ""
	if len(s) == 0 || len(s) < len(t) {
		return result
	}
	var freq [256]int
	for i := 0; i < len(t); i++ {
		freq[t[i]-'a']++
	}
	left, right, lastLeft, lastRight, count := 0, 0, -1, len(s), len(t)
	for ; right < len(s); right++ {
		if freq[s[right]-'a'] >= 1 {
			count--
		}
		freq[s[right]-'a']--
		for count == 0 {
			if right-left < lastRight-lastLeft {
				lastRight = right
				lastLeft = left
			}
			if freq[s[left]-'a'] == 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
	}
	if lastLeft != -1 {
		result = s[lastLeft : lastRight+1]
	}
	return result
}

func main() {

	S := "ADOBECODEBANC"
	T := "ABC"
	minWin := minWindow(S, T)
	fmt.Println(minWin)
}
