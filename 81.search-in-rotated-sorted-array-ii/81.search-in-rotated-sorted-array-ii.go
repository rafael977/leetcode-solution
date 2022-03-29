package main

/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}

	return false
}

// @lc code=end
