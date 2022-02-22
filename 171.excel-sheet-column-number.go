package main

/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	c := 0
	for _, r := range columnTitle {
		c = c*26 + int(r-'A'+1)
	}

	return c
}

// @lc code=end
