package leetcodesolution

/*
 * @lc app=leetcode id=1208 lang=golang
 *
 * [1208] Get Equal Substrings Within Budget
 */

// @lc code=start
func equalSubstring(s string, t string, maxCost int) (ans int) {
	costs := make([]int, len(s))
	for i := range s {
		c := int(s[i]) - int(t[i])
		if c < 0 {
			costs[i] = -c
		} else {
			costs[i] = c
		}
	}

	i, j := 0, 0
	tc := 0
	for j < len(s) {
		tc += costs[j]
		for tc > maxCost {
			tc -= costs[i]
			i++
		}
		ans = max(ans, j-i+1)
		j++
	}

	return
}

// @lc code=end
