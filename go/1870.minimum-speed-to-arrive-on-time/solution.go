package minimumspeedtoarriveontime

import "math"

/*
 * @lc app=leetcode id=1870 lang=golang
 *
 * [1870] Minimum Speed to Arrive on Time
 */

// @lc code=start
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if float64(n-1) >= hour {
		return -1
	}

	lo, hi := 1, 1
	for _, v := range dist {
		if hi < v {
			hi = v
		}
	}
	last := int(math.Ceil(float64(dist[n-1]) / (hour - float64(n-1))))
	if last > hi {
		hi = last
	}

	for lo < hi {
		x := lo + (hi-lo)/2
		var t float64 = 0
		for i, v := range dist {
			tv := float64(v) / float64(x)
			if i == n-1 {
				t += tv
			} else {
				t += math.Ceil(tv)
			}
		}

		if t > hour {
			lo = x + 1
		} else {
			hi = x
		}
	}
	return hi
}

// @lc code=end
