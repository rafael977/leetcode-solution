package combinationsumivq

import "sort"

/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if i < v {
				break
			}
			dp[i] += dp[i-v]
		}
	}
	return dp[target]
}

// @lc code=end
