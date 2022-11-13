package reversewordsinastring

import (
	"strings"
)

/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	fs := strings.Fields(s)
	sb := strings.Builder{}
	for i := len(fs) - 1; i >= 0; i-- {
		sb.WriteString(fs[i])
		if i != 0 {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}

// @lc code=end
