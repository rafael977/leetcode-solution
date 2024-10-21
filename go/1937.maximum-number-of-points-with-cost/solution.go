package maximumnumberofpointswithcost

import (
	"math"
	"slices"
)

/*
 * @lc app=leetcode id=1937 lang=golang
 *
 * [1937] Maximum Number of Points with Cost
 */

// @lc code=start
func maxPoints(points [][]int) int64 {
	n := len(points[0])
	dp := make([]int, n)
	for _, r := range points {
		newDp := make([]int, n)
		best := math.MinInt
		for i := 0; i < n; i++ {
			best = max(best, dp[i]+i)
			newDp[i] = max(newDp[i], best+r[i]-i)
		}
		best = math.MinInt
		for i := n - 1; i >= 0; i-- {
			best = max(best, dp[i]-i)
			newDp[i] = max(newDp[i], best+r[i]+i)
		}
		dp = newDp
	}

	return int64(slices.Max(dp))
}

// @lc code=end
