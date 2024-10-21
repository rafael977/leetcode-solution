package leetcodesolution

import "math"

/*
 * @lc app=leetcode id=3068 lang=golang
 *
 * [3068] Find the Maximum Sum of Node Values
 */

// @lc code=start
func maximumValueSum(nums []int, k int, _ [][]int) int64 {
	fe, fo := 0, math.MinInt64
	for _, v := range nums {
		fe, fo = max(fe+v, fo+(v^k)), max(fo+v, fe+(v^k))
	}

	return int64(fe)
}

// @lc code=end
