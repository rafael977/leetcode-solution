package triangle

import "math"

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	sums := make([][]int, 2)
	for i := range sums {
		sums[i] = make([]int, n)
	}

	sums[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		sc, sp := i%2, (i-1)%2
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				sums[sc][j] = triangle[i][j] + sums[sp][j]
			} else if j == len(triangle[i])-1 {
				sums[sc][j] = triangle[i][j] + sums[sp][j-1]
			} else {
				sums[sc][j] = triangle[i][j] + min(sums[sp][j-1], sums[sp][j])
			}
		}
	}

	i := (n - 1) % 2
	ans := math.MaxInt64
	for j := range sums[i] {
		ans = min(ans, sums[i][j])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
