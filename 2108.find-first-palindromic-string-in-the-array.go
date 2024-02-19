package leetcodesolution

/*
 * @lc app=leetcode id=2108 lang=golang
 *
 * [2108] Find First Palindromic String in the Array
 */

// @lc code=start
func firstPalindrome(words []string) string {
	isPalindrome := func(s string) bool {
		for i, j := 0, len(s)-1; i <= j; {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	for _, w := range words {
		if isPalindrome(w) {
			return w
		}
	}
	return ""
}

// @lc code=end
