package findplayerswithzerooronelosses

import "sort"

/*
 * @lc app=leetcode id=2225 lang=golang
 *
 * [2225] Find Players With Zero or One Losses
 */

// @lc code=start
func findWinners(matches [][]int) [][]int {
    losses := make(map[int]int)
	for _, v := range matches {
		if _, ok := losses[v[0]]; !ok {
			losses[v[0]] = 0
		}
		losses[v[1]]++
	}
	a1, a2 := make([]int, 0), make([]int, 0)
	for k, v := range losses {
		if v == 0 {
			a1 = append(a1, k)
		} else if v == 1 {
			a2 = append(a2, k)
		}
	}
	sort.Ints(a1)
	sort.Ints(a2)
	return [][]int{a1, a2}
}
// @lc code=end

