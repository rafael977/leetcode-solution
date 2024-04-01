package leetcodesolution

import "strings"

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	fs := strings.Fields(s)
	return len(fs[len(fs)-1])
}

// @lc code=end
