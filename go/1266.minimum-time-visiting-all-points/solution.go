package minimumtimevisitingallpoints

/*
 * @lc app=leetcode id=1266 lang=golang
 *
 * [1266] Minimum Time Visiting All Points
 */

// @lc code=start
func minTimeToVisitAllPoints(points [][]int) (ans int) {
	for i := 1; i < len(points); i++ {
		p1, p2 := points[i-1], points[i]
		x, y := abs(p1[0]-p2[0]), abs(p1[1]-p2[1])
		ans += max(x, y)
	}
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
