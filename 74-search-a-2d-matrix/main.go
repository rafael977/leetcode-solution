package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, h := 0, m*n-1
	for l <= h {
		p := l + (h-l)/2
		i, j := p/n, p%n
		if target == matrix[i][j] {
			return true
		}

		if matrix[i][j] < target {
			l = p + 1
		} else {
			h = p - 1
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}
