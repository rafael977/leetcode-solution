package leetcodesolution

import "slices"

/*
 * @lc app=leetcode id=2370 lang=golang
 *
 * [2370] Longest Ideal Subsequence
 */

// @lc code=start
func longestIdealString(s string, k int) (ans int) {
	dp := [26]int{}
	for _, c := range s {
		x := int(c - 'a')
		dp[x] = 1 + slices.Max(dp[max(0, x-k):min(26, x+k+1)])
	}
	return slices.Max(dp[:])
}

// @lc code=end
