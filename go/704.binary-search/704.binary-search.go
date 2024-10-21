package main

/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		p := (hi + lo) / 2
		if nums[p] == target {
			return p
		}
		if nums[p] < target {
			lo = p + 1
		} else {
			hi = p - 1
		}
	}
	return -1
}

// @lc code=end
