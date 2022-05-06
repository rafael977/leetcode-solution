package shortestunsortedcontinuoussubarray

import "math"

/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	min, max := math.MaxInt64, math.MinInt64
	l, r := -1, -1
	for i := 0; i < len(nums); i++ {
		if max > nums[i] {
			r = i
		} else {
			max = nums[i]
		}
		if min < nums[len(nums)-i-1] {
			l = len(nums) - i - 1
		} else {
			min = nums[len(nums)-i-1]
		}
	}
	if r == -1 {
		return 0
	}

	return r - l + 1
}

// @lc code=end
