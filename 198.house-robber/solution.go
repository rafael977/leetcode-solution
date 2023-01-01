package houserobber

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		best := 0
		for j := 0; j < i-1; j++ {
			best = max(best, dp[j])
		}
		dp[i] = best + nums[i]
	}

	return max(dp[n-1], dp[n-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
