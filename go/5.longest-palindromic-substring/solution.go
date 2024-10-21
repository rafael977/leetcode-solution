package longestpalindromicsubstring

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) (ans string) {
	maxL := 0
	for i := range s {
		ss := palindromLen(s, i, i)
		if len(ss) > maxL {
			maxL = len(ss)
			ans = ss
		}
		ss = palindromLen(s, i, i+1)
		if len(ss) > maxL {
			maxL = len(ss)
			ans = ss
		}
	}
	return
}

func palindromLen(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}

// @lc code=end
