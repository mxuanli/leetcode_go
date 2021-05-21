package main

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	// 最长公共子序列
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dpI := make([]int, n+1)
		dp[i] = dpI
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
				if dp[i-1][j] > dp[i][j] {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[m][n]
}
