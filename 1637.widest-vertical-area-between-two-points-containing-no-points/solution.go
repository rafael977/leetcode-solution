package widestverticalareabetweentwopointscontainingnopoints

import "sort"

/*
 * @lc app=leetcode id=1637 lang=golang
 *
 * [1637] Widest Vertical Area Between Two Points Containing No Points
 */

// @lc code=start
func maxWidthOfVerticalArea(points [][]int) (ans int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	x := points[0][0]
	for i := 1; i < len(points); i++ {
		px := points[i][0]
		if x < px {
			ans = max(ans, px-x)
			x = px
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
