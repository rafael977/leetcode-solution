package leetcodesolution

/*
 * @lc app=leetcode id=1614 lang=golang
 *
 * [1614] Maximum Nesting Depth of the Parentheses
 */

// @lc code=start
func maxDepth(s string) (ans int) {
	pc := 0
	for _, v := range s {
		if v == '(' {
			pc++
		} else if v == ')' {
			ans = max(ans, pc)
			pc--
		}
	}
	return
}

// @lc code=end
