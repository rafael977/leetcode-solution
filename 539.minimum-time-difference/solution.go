package minimumtimedifference

import (
	"sort"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=539 lang=golang
 *
 * [539] Minimum Time Difference
 */

// @lc code=start
func findMinDifference(timePoints []string) (ans int) {
	inMinutes := func(time string) int {
		st := strings.Split(time, ":")
		h, _ := strconv.Atoi(st[0])
		m, _ := strconv.Atoi(st[1])
		return h*60 + m
	}

	ans = 24 * 60
	sort.Strings(timePoints)
	for i := 0; i < len(timePoints)-1; i++ {
		ans = min(ans, inMinutes(timePoints[i+1])-inMinutes(timePoints[i]))
		if ans == 0 {
			return
		}
	}
	ans = min(ans, inMinutes(timePoints[0])+24*60-inMinutes(timePoints[len(timePoints)-1]))
	return
}

// @lc code=end
