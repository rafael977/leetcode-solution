package leetcodesolution

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 */

// @lc code=start
func findRelativeRanks(score []int) []string {
	si := make([][2]int, 0, len(score))
	for i, v := range score {
		si = append(si, [2]int{v, i})
	}
	sort.Slice(si, func(i, j int) bool {
		return si[i][0] > si[j][0]
	})

	ans := make([]string, len(score))
	for i, v := range si {
		if i == 0 {
			ans[v[1]] = "Gold Medal"
		} else if i == 1 {
			ans[v[1]] = "Silver Medal"
		} else if i == 2 {
			ans[v[1]] = "Bronze Medal"
		} else {
			ans[v[1]] = fmt.Sprintf("%d", i+1)
		}
	}
	return ans
}

// @lc code=end
