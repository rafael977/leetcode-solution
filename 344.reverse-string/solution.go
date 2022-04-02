package reversestring

/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

// @lc code=end
