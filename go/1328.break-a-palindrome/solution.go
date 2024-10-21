package breakapalindrome

/*
 * @lc app=leetcode id=1328 lang=golang
 *
 * [1328] Break a Palindrome
 */

// @lc code=start
func breakPalindrome(palindrome string) string {
	if len(palindrome) <= 1 {
		return ""
	}

	pb := []rune(palindrome)
	for i := 0; i < len(palindrome)/2; i++ {
		if pb[i] > 'a' {
			pb[i] = 'a'
			return string(pb)
		}
	}

	pb[len(pb)-1] = 'b'
	return string(pb)
}

// @lc code=end
