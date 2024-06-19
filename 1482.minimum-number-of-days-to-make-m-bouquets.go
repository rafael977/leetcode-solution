package leetcodesolution

import (
	"slices"
	"sort"
)

/*
 * @lc app=leetcode id=1482 lang=golang
 *
 * [1482] Minimum Number of Days to Make m Bouquets
 */

// @lc code=start
func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}

	max := slices.Max(bloomDay)
	ableToMake := func(x int) (r int) {
		cnt := 0
		for _, v := range bloomDay {
			if v <= x {
				cnt++
			} else {
				r += cnt / k
				cnt = 0
			}
		}
		r += cnt / k
		return
	}
	x := sort.Search(max, func(x int) bool {
		return ableToMake(x) >= m
	})

	return x
}

// @lc code=end
