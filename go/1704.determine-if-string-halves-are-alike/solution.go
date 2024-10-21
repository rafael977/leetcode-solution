package determineifstringhalvesarealike

import "strings"

/*
 * @lc app=leetcode id=1704 lang=golang
 *
 * [1704] Determine if String Halves Are Alike
 */

// @lc code=start
func halvesAreAlike(s string) bool {
	s = strings.ToLower(s)
    half := len(s)/2
	count := 0
	for i := range s {
		if isVowel(s[i]) {
			if i < half {
				count++
			} else {
				count--
			}
		}
	}
	return count == 0
}

func isVowel (c byte) bool {
	return c == 'a' || c== 'e' || c== 'i' || c == 'o' || c=='u'
}
// @lc code=end

