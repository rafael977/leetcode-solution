package numberofstepstoreduceanumbertozero

/*
 * @lc app=leetcode id=1342 lang=golang
 *
 * [1342] Number of Steps to Reduce a Number to Zero
 */

// @lc code=start
func numberOfSteps(num int) (steps int) {
	for num > 0 {
		if num&1 == 0 {
			num = num >> 1
		} else {
			num--
		}
		steps++
	}
	return
}

// @lc code=end
