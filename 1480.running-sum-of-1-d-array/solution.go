package runningsumof1darray

/*
 * @lc app=leetcode id=1480 lang=golang
 *
 * [1480] Running Sum of 1d Array
 */

// @lc code=start
func runningSum(nums []int) (rs []int) {
	rs = make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			rs[i] = nums[i]
			continue
		}
		rs[i] = rs[i-1] + nums[i]
	}
	return
}

// @lc code=end
