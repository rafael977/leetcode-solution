package findoriginalarrayfromdoubledarray

import "sort"

/*
 * @lc app=leetcode id=2007 lang=golang
 *
 * [2007] Find Original Array From Doubled Array
 */

// @lc code=start
func findOriginalArray(changed []int) []int {
	ans := make([]int, 0)

	sort.Ints(changed)
	double := make(map[int]int)
	for _, v := range changed {
		if v%2 == 1 || double[v] == 0 {
			double[v*2]++
			continue
		}

		ans = append(ans, v/2)
		double[v]--
		if double[v] == 0 {
			delete(double, v)
		}
	}

	if len(double) == 0 {
		return ans
	}
	return []int{}
}

// @lc code=end
