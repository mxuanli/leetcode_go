package main

import "fmt"

func partition(s string) [][]string {
	var res [][]string
	var path []string
	backtrack(s, &res, path)
	return res
}

func backtrack(s string, res *[][]string, path []string) {
	if len(s) == 0 {
		p := make([]string, len(path))
		copy(p, path)
		*res = append(*res, p)
		return
	}
	for i := 1; i < len(s)+1; i++ {
		if isPalindrome(s[:i]) {
			path := append(path, s[:i])
			backtrack(s[i:], res, path)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	s := "ababbbabbaba"
	r := partition(s)
	fmt.Println(r)
}
