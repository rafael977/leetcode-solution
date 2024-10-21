package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=1509 lang=golang
 *
 * [1509] Minimum Difference Between Largest and Smallest Value in Three Moves
 */

// @lc code=start
func minDifference(nums []int) (ans int) {
	n := len(nums)
	if n <= 3 {
		return 0
	}
	sort.Ints(nums)
	ans = nums[n-1] - nums[3]
	ans = min(ans, nums[n-2]-nums[2])
	ans = min(ans, nums[n-3]-nums[1])
	ans = min(ans, nums[n-4]-nums[0])
	return
}

// @lc code=end
