package leetcodesolution

import "math"

/*
 * @lc app=leetcode id=2485 lang=golang
 *
 * [2485] Find the Pivot Integer
 */

// @lc code=start
func pivotInteger(n int) int {
	m := (n*n + n) / 2
	x := int(math.Sqrt(float64(m)))
	if x*x == m {
		return x
	}
	return -1
}

// @lc code=end
