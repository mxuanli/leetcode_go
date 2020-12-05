package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack)+1-len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j+1 == len(needle) {
				return i
			}
		}
	}
	return -1
}

func main() {
	haystack := "abcdefg"
	needle := "g"
	r := strStr(haystack, needle)
	fmt.Println(r)
}
