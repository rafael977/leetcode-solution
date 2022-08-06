package kthsmallestelementinasortedmatrix

/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	check := func(mid int) bool {
		x := 0
		for i, j := n-1, 0; i >= 0 && j < n; {
			if matrix[i][j] <= mid {
				j++
				x += i + 1
			} else {
				i--
			}
		}
		return x >= k
	}

	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		mid := l + (r-l)/2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// @lc code=end
