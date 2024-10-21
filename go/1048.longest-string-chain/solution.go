package longeststringchain

import "sort"

/*
 * @lc app=leetcode id=1048 lang=golang
 *
 * [1048] Longest String Chain
 */

// @lc code=start
func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	dp := make([]int, len(words))
	for i := range dp {
		dp[i] = 1
	}
	for i := range words {
		for j := 0; j < i; j++ {
			if isPredecessor(words[j], words[i]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	max := 0
	for i := range dp {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func isPredecessor(a, b string) bool {
	if len(a)+1 != len(b) {
		return false
	}
	cFound := false
	for i, j := 0, 0; i < len(a) && j < len(b); {
		if a[i] != b[j] {
			if cFound {
				return false
			}
			j++
			cFound = true
		} else {
			i++
			j++
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
