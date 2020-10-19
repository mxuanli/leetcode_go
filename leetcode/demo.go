package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	left, right, maxLen := 0, -1, 0
	for left < len(s) {
		// 没有重复字符就移动右边界，有重复字符就移动左边界，把重复字符移动出去，直到左边界移动到最后
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main() {

	a := "pwwkew"
	maxLen := lengthOfLongestSubstring(a)
	fmt.Println(maxLen)
}
