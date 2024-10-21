package longestincreasingpathinamatrix

import "sort"

/*
 * @lc app=leetcode id=329 lang=golang
 *
 * [329] Longest Increasing Path in a Matrix
 */

// @lc code=start
type cell struct {
	x, y int
}

func longestIncreasingPath(matrix [][]int) (ans int) {
	m, n := len(matrix), len(matrix[0])
	dirs := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	cs := make([]cell, 0)
	for i := range matrix {
		for j := range matrix[i] {
			cs = append(cs, cell{x: i, y: j})
		}
	}

	sort.Slice(cs, func(i, j int) bool {
		return matrix[cs[i].x][cs[i].y] < matrix[cs[j].x][cs[j].y]
	})

	paths := make([][]int, m)
	for i := range paths {
		paths[i] = make([]int, n)
	}

	for _, c := range cs {
		paths[c.x][c.y] = 1
		for _, dir := range dirs {
			nx, ny := c.x+dir[0], c.y+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n ||
				matrix[nx][ny] >= matrix[c.x][c.y] {
				continue
			}

			paths[c.x][c.y] = max(paths[c.x][c.y], paths[nx][ny]+1)
		}
	}

	for i := range paths {
		for j := range paths[i] {
			ans = max(ans, paths[i][j])
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
