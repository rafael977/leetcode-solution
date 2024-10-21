package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=2971 lang=golang
 *
 * [2971] Find Polygon With the Largest Perimeter
 */

// @lc code=start
func largestPerimeter(nums []int) int64 {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	sum := 0
	for _, v := range nums {
		sum += v
	}
	for _, v := range nums {
		sum -= v
		if v < sum {
			return int64(sum + v)
		}
	}
	return -1
}

// @lc code=end
