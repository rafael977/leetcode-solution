package wordpattern

import "strings"

/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	sf := strings.Fields(s)
	c2w := make(map[rune]string)
	w2c := make(map[string]rune)

	if len(pattern) != len(sf) {
		return false
	}
	for i, c := range pattern {
		w := sf[i]
		if wfm, ok_c2w := c2w[c]; !ok_c2w {
			if _, ok := w2c[w]; !ok {
				c2w[c] = w
				w2c[w] = c
			} else {
				return false
			}
		} else if w != wfm {
			return false
		}
	}

	return true
}

// @lc code=end
