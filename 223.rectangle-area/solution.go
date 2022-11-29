package rectanglearea

/*
 * @lc app=leetcode id=223 lang=golang
 *
 * [223] Rectangle Area
 */

// @lc code=start
func computeArea(ax1 int, ay1 int,
	ax2 int, ay2 int,
	bx1 int, by1 int,
	bx2 int, by2 int) int {
	l, b, r, t := max(ax1, bx1), max(ay1, by1), min(ax2, bx2), min(ay2, by2)

	overlap := 0
	if l < r && b < t {
		overlap = (r - l) * (t - b)
	}

	return (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1) - overlap
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
