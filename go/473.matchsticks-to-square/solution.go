package matchstickstosquare

import "sort"

/*
 * @lc app=leetcode id=473 lang=golang
 *
 * [473] Matchsticks to Square
 */

// @lc code=start
func makesquare(matchsticks []int) bool {
	total := 0
	for _, v := range matchsticks {
		total += v
	}
	if total%4 != 0 {
		return false
	}

	eLen := total / 4
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	edge := make([]int, 4)
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i == len(matchsticks) {
			return true
		}

		for e := range edge {
			edge[e] += matchsticks[i]
			if edge[e] <= eLen && dfs(i+1) {
				return true
			}
			edge[e] -= matchsticks[i]
		}
		return false
	}

	return dfs(0)
}

// @lc code=end
