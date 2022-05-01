package backspacestringcompare

/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	genString := func(s string) string {
		stack := make([]rune, 0, len(s))
		for _, c := range s {
			if c == '#' {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, c)
			}
		}
		return string(stack)
	}

	return genString(s) == genString(t)
}

// @lc code=end
