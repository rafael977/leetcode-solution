package main

/*
 * @lc app=leetcode id=413 lang=golang
 *
 * [413] Arithmetic Slices
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) (ans int) {
	calcSub := func(n int) (c int) {
		for i := n - 2; i > 0; i-- {
			c += i
		}
		return
	}

	if len(nums) < 3 {
		return 0
	}
	diff := nums[1] - nums[0]
	for i, j := 0, 1; j <= len(nums); j++ {
		if j == len(nums) || nums[j]-nums[j-1] != diff {
			ans += calcSub(j - i)
			i = j - 1
			if j < len(nums) {
				diff = nums[j] - nums[j-1]
			}
		}
	}
	return
}

// @lc code=end
