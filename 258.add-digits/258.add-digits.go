package main

/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

// @lc code=start
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	r := num % 9
	if r == 0 {
		return 9
	}
	return r
}

// @lc code=end
