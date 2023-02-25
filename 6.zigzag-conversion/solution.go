package zigzagconversion

import "strings"

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	rows := make([][]byte, numRows)
	for i := range rows {
		rows[i] = make([]byte, 0)
	}

	for r, inc, i := 0, -1, 0; i < len(s); i++ {
		rows[r] = append(rows[r], s[i])
		if r == 0 || r == numRows-1 {
			inc = -inc
		}
		r += inc
	}
	sb := strings.Builder{}
	for _, r := range rows {
		sb.WriteString(string(r))
	}

	return sb.String()
}

// @lc code=end
