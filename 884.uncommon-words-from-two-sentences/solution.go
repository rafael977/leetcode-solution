package uncommonwordsfromtwosentences

import "strings"

/*
 * @lc app=leetcode id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 */

// @lc code=start
func uncommonFromSentences(s1 string, s2 string) (ans []string) {
	freq := make(map[string]int)
	for _, v := range strings.Fields(s1) {
		freq[v]++
	}
	for _, v := range strings.Fields(s2) {
		freq[v]++
	}

	for k, v := range freq {
		if v == 1 {
			ans = append(ans, k)
		}
	}
	return
}

// @lc code=end
