package rangesumquery2dimmutable

/*
 * @lc app=leetcode id=304 lang=golang
 *
 * [304] Range Sum Query 2D - Immutable
 */

// @lc code=start
type NumMatrix struct {
	sm [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sm := make([][]int, m)
	for i := range sm {
		sm[i] = make([]int, n)
		for j := range sm[i] {
			sm[i][j] = matrix[i][j]
			if i >= 1 {
				sm[i][j] += sm[i-1][j]
			}
			if j >= 1 {
				sm[i][j] += sm[i][j-1]
			}
			if i >= 1 && j >= 1 {
				sm[i][j] -= sm[i-1][j-1]
			}
		}
	}

	return NumMatrix{sm: sm}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (sum int) {
	sum = this.sm[row2][col2]
	if row1 >= 1 {
		sum -= this.sm[row1-1][col2]
	}
	if col1 >= 1 {
		sum -= this.sm[row2][col1-1]
	}
	if row1 >= 1 && col1 >= 1 {
		sum += this.sm[row1-1][col1-1]
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end
