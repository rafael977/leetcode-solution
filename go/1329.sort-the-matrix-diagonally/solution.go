package sortthematrixdiagonally

import "sort"

/*
 * @lc app=leetcode id=1329 lang=golang
 *
 * [1329] Sort the Matrix Diagonally
 */

// @lc code=start
func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		arr := make([]int, 0)
		for x, y := i, 0; x < m && y < n; x, y = x+1, y+1 {
			arr = append(arr, mat[x][y])
		}
		sort.Ints(arr)
		for x, y := i, 0; x < m && y < n; x, y = x+1, y+1 {
			mat[x][y] = arr[y]
		}
	}
	for j := 1; j < n; j++ {
		arr := make([]int, 0)
		for x, y := 0, j; x < m && y < n; x, y = x+1, y+1 {
			arr = append(arr, mat[x][y])
		}
		sort.Ints(arr)
		for x, y := 0, j; x < m && y < n; x, y = x+1, y+1 {
			mat[x][y] = arr[x]
		}
	}
	return mat
}

// @lc code=end
