package main

import "sort"

/*
 * @lc app=leetcode id=1029 lang=golang
 *
 * [1029] Two City Scheduling
 */

// @lc code=start
func twoCitySchedCost(costs [][]int) (ans int) {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

	for i := 0; i < len(costs)/2; i++ {
		ans += costs[i][0]
	}
	for i := len(costs) / 2; i < len(costs); i++ {
		ans += costs[i][1]
	}

	return ans
}

// @lc code=end
