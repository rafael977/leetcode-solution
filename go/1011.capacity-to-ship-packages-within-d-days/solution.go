package capacitytoshippackageswithinddays

/*
 * @lc app=leetcode id=1011 lang=golang
 *
 * [1011] Capacity To Ship Packages Within D Days
 */

// @lc code=start
func shipWithinDays(weights []int, days int) int {
	lo, hi := 0, 0
	for _, w := range weights {
		if w > lo {
			lo = w
		}
		hi += w
	}

	for lo < hi {
		mid := lo + (hi-lo)/2
		d := 1
		cur := 0
		for _, w := range weights {
			if cur+w > mid {
				d++
				cur = 0
			}
			cur += w
		}

		if d <= days {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

// @lc code=end
