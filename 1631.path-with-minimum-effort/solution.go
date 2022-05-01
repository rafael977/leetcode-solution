package pathwithminimumeffort

import "sort"

/*
 * @lc app=leetcode id=1631 lang=golang
 *
 * [1631] Path With Minimum Effort
 */

// @lc code=start
type edge struct {
	u, v, w int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func find(parents map[int]int, v int) int {
	if parents[v] != v {
		parents[v] = find(parents, parents[v])
	}
	return parents[v]
}

func union(parents map[int]int, x, y int) {
	px, py := find(parents, x), find(parents, y)
	if px < py {
		parents[py] = px
	} else {
		parents[px] = py
	}
}

func minimumEffortPath(heights [][]int) int {
	row, col := len(heights), len(heights[0])
	edges := make([]edge, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			pos := i*col + j
			if i < row-1 {
				edges = append(edges, edge{u: pos, v: pos + col, w: abs(heights[i][j] - heights[i+1][j])})
			}
			if j < col-1 {
				edges = append(edges, edge{u: pos, v: pos + 1, w: abs(heights[i][j] - heights[i][j+1])})
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	parents := make(map[int]int)
	for i := 0; i < row*col; i++ {
		parents[i] = i
	}

	max := 0
	e := 0
	for find(parents, 0) != find(parents, row*col-1) {
		edge := edges[e]
		union(parents, edge.u, edge.v)
		max = edge.w
		e++
	}

	return max
}

// @lc code=end
