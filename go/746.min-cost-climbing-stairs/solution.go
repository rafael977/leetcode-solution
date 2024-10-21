package mincostclimbingstairs

/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	f := make([]int, n)
	f[0] = cost[0]
	f[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		f[i] = min(f[i-2], f[i-1]) + cost[i]
	}
	return min(f[n-1], f[n-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
