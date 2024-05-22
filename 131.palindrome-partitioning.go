package leetcodesolution

/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) (ans [][]string) {
	isPalindrome := func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	var backtrack func(i, j int, ss []string)
	backtrack = func(i, j int, ss []string) {
		if i == len(s) {
			ans = append(ans, ss)
			return
		}
		if j == len(s) {
			return
		}
		if j < len(s) && isPalindrome(i, j) {
			backtrack(j+1, j+1, append(append([]string{}, ss...), s[i:j+1]))
		}
		backtrack(i, j+1, ss)
	}
	backtrack(0, 0, []string{})
	return
}

// @lc code=end
