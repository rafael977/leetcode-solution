package checkiftwostringarraysareequivalent

import "strings"

/*
 * @lc app=leetcode id=1662 lang=golang
 *
 * [1662] Check If Two String Arrays are Equivalent
 */

// @lc code=start
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	sb1, sb2 := strings.Builder{}, strings.Builder{}
	for _, v := range word1 {
		sb1.WriteString(v)
	}
	for _, v := range word2 {
		sb2.WriteString(v)
	}
	return sb1.String() == sb2.String()
}

// @lc code=end
