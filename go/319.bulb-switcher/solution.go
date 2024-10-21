package bulbswitcher

import "math"

/*
 * @lc app=leetcode id=319 lang=golang
 *
 * [319] Bulb Switcher
 */

// @lc code=start

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}

// @lc code=end
