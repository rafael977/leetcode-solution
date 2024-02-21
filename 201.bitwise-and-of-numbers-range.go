package leetcodesolution

/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	ans := left
	for i := left; i <= right; i++ {
		ans &= i
		if ans == 0 {
			break
		}
	}
	return ans
}

// @lc code=end
