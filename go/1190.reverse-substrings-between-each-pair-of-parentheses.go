package leetcodesolution

/*
 * @lc app=leetcode id=1190 lang=golang
 *
 * [1190] Reverse Substrings Between Each Pair of Parentheses
 */

// @lc code=start
func reverseParentheses(s string) string {
	st := make([]rune, 0, len(s))
	for _, v := range s {
		if v == ')' {
			tmp := make([]rune, 0, len(st))
			i := len(st) - 1
			for i >= 0 && st[i] != '(' {
				tmp = append(tmp, st[i])
				i--
			}
			st = append(st[:i], tmp...)
			continue
		}
		st = append(st, v)
	}

	return string(st)
}

// @lc code=end
