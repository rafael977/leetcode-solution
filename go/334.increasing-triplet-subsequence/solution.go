package increasingtripletsubsequence

import "math"

/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for i := 1; i < n; i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}
	return false
}

// @lc code=end
