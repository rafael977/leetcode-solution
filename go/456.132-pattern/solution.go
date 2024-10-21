package pattern132

import "math"

/*
 * @lc app=leetcode id=456 lang=golang
 *
 * [456] 132 Pattern
 */

// @lc code=start
func find132pattern(nums []int) bool {
	nk := math.MinInt64
	stack := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < nk {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			nk = max(nk, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
