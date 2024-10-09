package minimumaddtomakeparenthesesvalid

/*
 * @lc app=leetcode id=921 lang=golang
 *
 * [921] Minimum Add to Make Parentheses Valid
 */

// @lc code=start
func minAddToMakeValid(s string) (ans int) {
	c := 0
	for _, v := range s {
		if v == '(' {
			c++
		} else if c > 0 {
			c--
		} else {
			ans++
		}
	}
	return ans + c
}

// @lc code=end
