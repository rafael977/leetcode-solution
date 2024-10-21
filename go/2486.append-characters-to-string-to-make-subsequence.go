package leetcodesolution

/*
 * @lc app=leetcode id=2486 lang=golang
 *
 * [2486] Append Characters to String to Make Subsequence
 */

// @lc code=start
func appendCharacters(s string, t string) int {
	j := 0
	for i := range s {
		if s[i] == t[j] {
			j++
		}
		if j == len(t) {
			return 0
		}
	}
	return len(t) - j
}

// @lc code=end
