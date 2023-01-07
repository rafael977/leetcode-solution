package detectcapital

import (
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	upper, lower := strings.ToUpper(word), strings.ToLower(word)
	lr := []rune(lower)
	lr[0] = unicode.ToUpper(lr[0])
	fUpper := string(lr)
	return word == upper || word == lower || word == fUpper
}

// @lc code=end
