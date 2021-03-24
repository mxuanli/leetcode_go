package main

import "fmt"

func numDistinct(s string, t string) int {
	n := len(s)
	m := len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

func main() {
	s := "babgbag"
	t := "bag"
	r := numDistinct(s, t)
	fmt.Println(r)
}
