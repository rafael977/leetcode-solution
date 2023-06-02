package matrixdiagonalsum

/*
 * @lc app=leetcode id=1572 lang=golang
 *
 * [1572] Matrix Diagonal Sum
 */

// @lc code=start
func diagonalSum(mat [][]int) (ans int) {
	n := len(mat)
	for i := 0; i < n; i++ {
		ans += mat[i][i]
		if i != n-1-i {
			ans += mat[i][n-1-i]
		}
	}
	return
}

// @lc code=end
