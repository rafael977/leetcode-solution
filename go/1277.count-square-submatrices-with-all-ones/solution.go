// Created by Rafael Shen at 2024/10/27 10:55
// leetgo: 1.4.9
// https://leetcode.com/problems/count-square-submatrices-with-all-ones/

package countsquaresubmatriceswithallones

// @lc code=begin

func countSquares(matrix [][]int) (ans int) {
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				f[i][j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				f[i][j] = 0
			} else {
				f[i][j] = 1 + min(min(f[i-1][j], f[i][j-1]), f[i-1][j-1])
			}

			ans += f[i][j]
		}
	}
	return
}

// @lc code=end
