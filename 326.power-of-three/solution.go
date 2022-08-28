package powerofthree

/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}

// @lc code=end
