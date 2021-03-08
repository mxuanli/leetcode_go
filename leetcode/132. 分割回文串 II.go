package main

import "fmt"

func minCut(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = n
	}
	for i := 0; i < n; i++ {
		if isPalindrome(s[:i+1]) {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if isPalindrome(s[j+1 : i+1]) {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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
	s := "cdd"
	r := minCut(s)
	fmt.Println(r)
}
