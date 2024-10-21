package greatestcommondivisorofstrings

import "strings"

/*
 * @lc app=leetcode id=1071 lang=golang
 *
 * [1071] Greatest Common Divisor of Strings
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	l := l1
	if l > l2 {
		l = l2
	}
	for i := l; i > 0; i-- {
		if l1%i == 0 && l2%i == 0 {
			d := str1[:i]
			if check(str1, d) && check(str2, d) {
				return d
			}
		}
	}
	return ""
}

func check(s, d string) bool {
	n := len(s) / len(d)
	ss := strings.Builder{}
	for i := 0; i < n; i++ {
		ss.WriteString(d)
	}
	return s == ss.String()
}

// @lc code=end
