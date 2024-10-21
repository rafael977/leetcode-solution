package palindromicsubstrings

/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */

// @lc code=start
func countSubstrings(s string) (ans int) {
	countPalindromWithCenter := func(l, r int) int {
		c := 0
		for l >= 0 && r < len(s) &&
			s[l] == s[r] {
			c++
			l--
			r++
		}
		return c
	}

	for i := range s {
		ans += countPalindromWithCenter(i, i)
		if i != len(s)-1 {
			ans += countPalindromWithCenter(i, i+1)
		}
	}

	return
}

// @lc code=end
