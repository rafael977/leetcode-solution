package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=1608 lang=golang
 *
 * [1608] Special Array With X Elements Greater Than or Equal X
 */

// @lc code=start
func specialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for x := 1; x <= len(nums); x++ {
		if (x == len(nums) && nums[x-1] >= x) ||
			(nums[x-1] >= x && nums[x] < x) {
			return x
		}
	}

	return -1
}

// @lc code=end
