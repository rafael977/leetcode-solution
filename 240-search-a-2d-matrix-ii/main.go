package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	l, r := 0, m-1
	row := 0
	for row < n {
		if l > r || matrix[row][r] < target {
			row++
			l = 0
			continue
		}
		mid := (l + r) / 2
		// fmt.Println(mid)
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(matrix, 5))
	fmt.Println(searchMatrix(matrix, 20))
}
