package insertinterval

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	l, r := newInterval[0], newInterval[1]
	merged := false
	for _, v := range intervals {
		if r < v[0] {
			if !merged {
				ans = append(ans, []int{l, r})
				merged = true
			}
			ans = append(ans, v)
		} else if l > v[1] {
			ans = append(ans, v)
		} else {
			l = min(l, v[0])
			r = max(r, v[1])
		}
	}
	if !merged {
		ans = append(ans, []int{l, r})
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
