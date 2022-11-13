package toeplitzmatrix

/*
 * @lc app=leetcode id=766 lang=golang
 *
 * [766] Toeplitz Matrix
 */

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if i <= j && matrix[i][j] != matrix[0][j-i] {
				return false
			}
			if i > j && matrix[i][j] != matrix[i-j][0] {
				return false
			}
		}
	}
	return true
}

// @lc code=end
