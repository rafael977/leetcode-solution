package sortarraybyparity

/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 */

// @lc code=start
func sortArrayByParity(nums []int) []int {
	slow := 0
	for fast := range nums {
		if nums[fast]%2 == 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	return nums
}

// @lc code=end
