package minimumnumberofarrowstoburstballoons

import "sort"

/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start
func findMinArrowShots(points [][]int) (ans int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	ans = 1
	right := points[0][1]
	for _, p := range points {
		if p[0] > right {
			ans++
			right = p[1]
		}
	}
	return
}

// @lc code=end
