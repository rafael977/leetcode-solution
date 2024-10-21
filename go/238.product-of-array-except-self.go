package leetcodesolution

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		ans[i] *= ans[i-1] * nums[i-1]
	}
	x := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			x = nums[i]
			continue
		}
		ans[i] *= x
		x *= nums[i]
	}

	return ans
}

// @lc code=end
