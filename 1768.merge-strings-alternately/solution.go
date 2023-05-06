package mergestringsalternately

import "strings"

/*
 * @lc app=leetcode id=1768 lang=golang
 *
 * [1768] Merge Strings Alternately
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	n := len(word1)
	if n > len(word2) {
		n = len(word2)
	}
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[i])
	}
	sb.WriteString(word1[n:])
	sb.WriteString(word2[n:])
	return sb.String()
}

// @lc code=end
