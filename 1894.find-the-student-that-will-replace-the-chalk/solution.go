package findthestudentthatwillreplacethechalk

/*
 * @lc app=leetcode id=1894 lang=golang
 *
 * [1894] Find the Student that Will Replace the Chalk
 */

// @lc code=start
func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, v := range chalk {
		sum += v
	}
	r := k % sum
	for i, v := range chalk {
		if r < v {
			return i
		}
		r -= v
	}
	return 0
}

// @lc code=end
