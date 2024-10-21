package maximumproductdifferencebetweentwopairs

import "sort"

/*
 * @lc app=leetcode id=1913 lang=golang
 *
 * [1913] Maximum Product Difference Between Two Pairs
 */

// @lc code=start
func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n-1]*nums[n-2] - nums[1]*nums[0]
}

// @lc code=end
