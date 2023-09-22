package minimumnumberoftapstoopentowateragarden

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=1326 lang=golang
 *
 * [1326] Minimum Number of Taps to Open to Water a Garden
 */

// @lc code=start
func minTaps(n int, ranges []int) int {
	intervals := make([][2]int, 0)
	for i, v := range ranges {
		s, e := max(0, i-v), min(n, i+v)
		intervals = append(intervals, [2]int{s, e})
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for _, v := range intervals {
		s, e := v[0], v[1]
		if dp[s] == math.MaxInt {
			return -1
		}
		for i := s; i <= e; i++ {
			dp[i] = min(dp[i], dp[s]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
