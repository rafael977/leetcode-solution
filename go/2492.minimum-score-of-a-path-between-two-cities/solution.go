package minimumscoreofapathbetweentwocities

import "math"

/*
 * @lc app=leetcode id=2492 lang=golang
 *
 * [2492] Minimum Score of a Path Between Two Cities
 */

// @lc code=start
func minScore(n int, roads [][]int) (ans int) {
	p := make([]int, n+1)
	for i := range p {
		p[i] = i
	}

	for _, r := range roads {
		union(r[0], r[1], p)
	}

	ans = math.MaxInt
	for _, r := range roads {
		if find(r[0], p) == 1 {
			if ans > r[2] {
				ans = r[2]
			}
		}
	}
	return
}

func find(i int, parent []int) int {
	if parent[i] != i {
		parent[i] = find(parent[i], parent)
	}
	return parent[i]
}

func union(i, j int, parent []int) {
	pi, pj := find(i, parent), find(j, parent)
	if pi < pj {
		parent[pj] = pi
	} else {
		parent[pi] = pj
	}
}

// @lc code=end
