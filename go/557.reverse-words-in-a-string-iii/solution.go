package reversewordsinastringiii

import "strings"

/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
func reverseWords(s string) string {
	fs := strings.Fields(s)
	sb := strings.Builder{}
	for _, f := range fs {
		sb.WriteString(reverse(f))
		sb.WriteString(" ")
	}

	rs := sb.String()
	return rs[:len(rs)-1]
}

func reverse(s string) string {
	sb := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

// @lc code=end
