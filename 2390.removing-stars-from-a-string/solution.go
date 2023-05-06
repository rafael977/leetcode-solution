package removingstarsfromastring

/*
 * @lc app=leetcode id=2390 lang=golang
 *
 * [2390] Removing Stars From a String
 */

// @lc code=start
func removeStars(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		if s[i] != '*' {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

// @lc code=end
