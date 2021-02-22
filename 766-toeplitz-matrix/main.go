package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				continue
			}

			if i <= j && matrix[0][j-i] != matrix[i][j] {
				return false
			} else if i > j && matrix[i-j][0] != matrix[i][j] {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}))
	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
}
