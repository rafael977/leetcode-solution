package russiandollenvelopes

import "sort"

/*
 * @lc app=leetcode id=354 lang=golang
 *
 * [354] Russian Doll Envelopes
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1])
	})

	hs := make([]int, len(envelopes))
	for i := range envelopes {
		hs[i] = envelopes[i][1]
	}
	return lengthLIS(hs)
}

func lengthLIS(arr []int) int {
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = 1
	}

	for i := range arr {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	max := 0
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max
}

// @lc code=end
