package numberoflongestincreasingsubsequence

/*
 * @lc app=leetcode id=673 lang=golang
 *
 * [673] Number of Longest Increasing Subsequence
 */

// @lc code=start
func findNumberOfLIS(nums []int) (ans int) {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = []int{1, 1}
	}

	for i, v := range nums {
		for j := 0; j < i; j++ {
			if v <= nums[j] {
				continue
			}
			if dp[i][0] < dp[j][0]+1 {
				dp[i][0] = dp[j][0] + 1
				dp[i][1] = dp[j][1]
			} else if dp[i][0] == dp[j][0]+1 {
				dp[i][1] += dp[j][1]
			}
		}
	}
	max := 0
	for _, v := range dp {
		if max < v[0] {
			max = v[0]
		}
	}
	for _, v := range dp {
		if v[0] == max {
			ans += v[1]
		}
	}
	return
}

// @lc code=end
