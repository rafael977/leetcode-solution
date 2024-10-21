package onesandzeroes

import "strings"

/*
 * @lc app=leetcode id=474 lang=golang
 *
 * [474] Ones and Zeroes
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i, s := range strs {
		c0 := strings.Count(s, "0")
		c1 := len(s) - c0
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i+1][j][k] = dp[i][j][k]
				if j >= c0 && k >= c1 {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-c0][k-c1]+1)
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
