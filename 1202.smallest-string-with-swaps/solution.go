package smalleststringwithswaps

import "sort"

/*
 * @lc app=leetcode id=1202 lang=golang
 *
 * [1202] Smallest String With Swaps
 */

// @lc code=start
type subset struct {
	s  []rune
	ci int
}

func find(parents map[int]int, idx int) int {
	if parents[idx] != idx {
		parents[idx] = find(parents, parents[idx])
	}

	return parents[idx]
}

func union(parents map[int]int, x, y int) {
	px, py := find(parents, x), find(parents, y)
	if px < py {
		parents[py] = px
	} else {
		parents[px] = py
	}
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	rs := []rune(s)
	n := len(rs)

	parents := make(map[int]int)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

	for _, p := range pairs {
		union(parents, p[0], p[1])
	}

	subsets := make(map[int]subset)
	for i := 0; i < n; i++ {
		pi := find(parents, i)
		ss, ok := subsets[pi]
		if !ok {
			ss = subset{s: make([]rune, 0), ci: 0}
		}
		ss.s = append(ss.s, rs[i])
		subsets[pi] = ss
	}

	for _, ss := range subsets {
		sort.Slice(ss.s, func(i, j int) bool {
			return ss.s[i] < ss.s[j]
		})
	}

	for i := 0; i < n; i++ {
		pi := find(parents, i)
		ss := subsets[pi]
		rs[i] = ss.s[ss.ci]
		ss.ci++
		subsets[pi] = ss
	}
	return string(rs)
}

// @lc code=end
