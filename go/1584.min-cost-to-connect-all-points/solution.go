package mincosttoconnectallpoints

import (
	"sort"
)

/*
 * @lc app=leetcode id=1584 lang=golang
 *
 * [1584] Min Cost to Connect All Points
 */

// @lc code=start
type edge struct {
	u, v, w int
}

func weight(p1, p2 []int) int {
	wx := p1[0] - p2[0]
	if wx < 0 {
		wx = -wx
	}
	wy := p1[1] - p2[1]
	if wy < 0 {
		wy = -wy
	}
	return wx + wy
}

func find(parents map[int]int, v int) int {
	if parents[v] != v {
		parents[v] = find(parents, parents[v])
	}
	return parents[v]
}

func minCostConnectPoints(points [][]int) (ans int) {
	edges := make([]edge, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, edge{u: i, v: j, w: weight(points[i], points[j])})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	parents := make(map[int]int)
	for i := 0; i < len(points); i++ {
		parents[i] = i
	}

	for _, e := range edges {
		pu, pv := find(parents, e.u), find(parents, e.v)
		if pu != pv {
			ans += e.w
			parents[pu] = pv
		}
	}
	return
}

// @lc code=end
