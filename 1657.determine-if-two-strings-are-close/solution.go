package determineiftwostringsareclose

import "sort"

/*
 * @lc app=leetcode id=1657 lang=golang
 *
 * [1657] Determine if Two Strings Are Close
 */

// @lc code=start
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	c1, c2 := make([]int, 26), make([]int, 26)
	for i := range word1 {
		c1[word1[i]-'a']++
		c2[word2[i]-'a']++
	}
	for i := range c1 {
		if (c1[i] > 0 && c2[i] == 0) || (c2[i] > 0 && c1[i] == 0) {
			return false
		}
	}
	sort.Ints(c1)
	sort.Ints(c2)
	for i := range c1 {
		if c1[i] != c2[i] {
			return false
		}
	}
	return true
}

// @lc code=end
