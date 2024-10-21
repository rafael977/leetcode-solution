package searcha2dmatrix

/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// m, n := len(matrix), len(matrix[0])
	// l, h := 0, m*n-1
	// for l <= h {
	// 	p := l + (h-l)/2
	// 	i, j := p/n, p%n
	// 	if target == matrix[i][j] {
	// 		return true
	// 	}

	// 	if matrix[i][j] < target {
	// 		l = p + 1
	// 	} else {
	// 		h = p - 1
	// 	}
	// }

	// return false

	for _, row := range matrix {
		if row[0] > target {
			return false
		}
		if row[len(row)-1] < target {
			continue
		}

		i, j := 0, len(row)-1
		for i <= j {
			p := (i + j) / 2
			if row[p] == target {
				return true
			}
			if row[p] > target {
				j = p - 1
			} else {
				i = p + 1
			}
		}
	}

	return false
}

// @lc code=end
