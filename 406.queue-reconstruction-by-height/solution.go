package queuereconstructionbyheight

import (
	"sort"
)

/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		pi, pj := people[i], people[j]

		// samller k first
		if pi[1] < pj[1] {
			return true
		} else if pi[1] > pj[1] {
			return false
		} else {
			// same k, if k=0, smaller h first, else bigger h first
			if pi[1] == 0 {
				return pi[0] < pj[0]
			} else {
				return pi[0] > pj[0]
			}
		}
	})

	ans := make([][]int, len(people))
	for i, p := range people {
		if p[1] == 0 {
			ans[i] = p
			continue
		}

		k := p[1]
		x := 0
		for x < len(ans) {
			if ans[x][0] >= p[0] {
				k--
			}
			x++
			if k == 0 {
				break
			}
		}
		copy(ans[x+1:], ans[x:])
		ans[x] = p
	}

	return ans
}

// @lc code=end
