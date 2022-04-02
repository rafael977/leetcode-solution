package validpalindromeii

/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 */

// @lc code=start
func validPalindrome(s string) bool {
	palindrome := func(lo, hi int) bool {
		for i, j := lo, hi; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return palindrome(i+1, j) || palindrome(i, j-1)
		}
	}
	return true
}

// @lc code=end
