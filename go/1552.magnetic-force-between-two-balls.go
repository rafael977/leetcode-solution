package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=1552 lang=golang
 *
 * [1552] Magnetic Force Between Two Balls
 */

// @lc code=start
func maxDistance(position []int, m int) int {
	sort.Ints(position)
	countBalls := func(x int) int {
		c := 1
		l := position[0]
		for _, v := range position {
			if v-l >= x {
				c++
				l = v
			}
		}
		return c
	}
	return sort.Search(position[len(position)-1]-position[0], func(i int) bool {
		x := i + 1
		cb := countBalls(x)
		return cb < m
	})
}

// @lc code=end
