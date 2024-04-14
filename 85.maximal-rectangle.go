package leetcodesolution

/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) (ans int) {
	n, m := len(matrix), len(matrix[0])
	ps := make([][]int, n)
	for i := 0; i < n; i++ {
		ps[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i > 0 {
				ps[i][j] += ps[i-1][j]
			}
			if j > 0 {
				ps[i][j] += ps[i][j-1]
			}
			if i > 0 && j > 0 {
				ps[i][j] -= ps[i-1][j-1]
			}
			ps[i][j] += int(matrix[i][j] - '0')
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for ii := i; ii < n; ii++ {
				for jj := j; jj < m; jj++ {
					sum := ps[ii][jj]
					if i > 0 {
						sum -= ps[i-1][jj]
					}
					if j > 0 {
						sum -= ps[ii][j-1]
					}
					if i > 0 && j > 0 {
						sum += ps[i-1][j-1]
					}
					if sum == (ii-i+1)*(jj-j+1) {
						ans = max(ans, sum)
					}
				}
			}
		}
	}
	return
}

// @lc code=end
