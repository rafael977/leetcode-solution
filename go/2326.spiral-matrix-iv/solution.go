package spiralmatrixiv

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2326 lang=golang
 *
 * [2326] Spiral Matrix IV
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = 1001
		}
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d := 0
	p := [2]int{0, 0}
	for i := 0; i < m*n; i++ {
		if head == nil {
			ans[p[0]][p[1]] = -1
		} else {
			ans[p[0]][p[1]] = head.Val
			head = head.Next
		}
		np := [2]int{p[0] + dirs[d][0], p[1] + dirs[d][1]}
		if np[0] < 0 || np[0] == m ||
			np[1] < 0 || np[1] == n ||
			ans[np[0]][np[1]] != 1001 {
			d = (d + 1) % 4
			np = [2]int{p[0] + dirs[d][0], p[1] + dirs[d][1]}
		}
		p = np
	}

	return ans
}

// @lc code=end
