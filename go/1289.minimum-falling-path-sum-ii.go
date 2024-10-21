package leetcodesolution

import (
	"math"
	"slices"
)

/*
 * @lc app=leetcode id=1289 lang=golang
 *
 * [1289] Minimum Falling Path Sum II
 */

// @lc code=start
func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	if n == 1 {
		return grid[0][0]
	}

	fs, ss := math.MaxInt, math.MaxInt
	for _, v := range grid[0] {
		if v < fs {
			ss = fs
			fs = v
		} else if v < ss {
			ss = v
		}
	}
	for i := 1; i < n; i++ {
		nfs, nss := math.MaxInt, math.MaxInt
		for j, v := range grid[i] {
			if grid[i-1][j] == fs {
				grid[i][j] = v + ss
			} else {
				grid[i][j] = v + fs
			}
			if grid[i][j] < nfs {
				nss = nfs
				nfs = grid[i][j]
			} else if grid[i][j] < nss {
				nss = grid[i][j]
			}
		}
		fs, ss = nfs, nss
	}

	return slices.Min(grid[n-1])
}

// @lc code=end
