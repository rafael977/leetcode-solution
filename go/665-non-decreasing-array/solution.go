package nondecreasingarray

/*
 * @lc app=leetcode id=665 lang=golang
 *
 * [665] Non-decreasing Array
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	c := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			c--
			if c < 0 {
				return false
			}

			if i == 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				// need to increase nums[i] as both i-2 & i-1 are greater than i
				nums[i] = nums[i-1]
			}
		}
	}

	return true
}

// @lc code=end
