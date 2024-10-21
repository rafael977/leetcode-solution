package minimumpenaltyforashop

import "strings"

/*
 * @lc app=leetcode id=2483 lang=golang
 *
 * [2483] Minimum Penalty for a Shop
 */

// @lc code=start
func bestClosingTime(customers string) (ans int) {
	cost := strings.Count(customers, "Y")
	min := cost
	for i, v := range customers {
		if v == 'N' {
			cost++
		} else {
			cost--
			if cost < min {
				min = cost
				ans = i + 1
			}
		}
	}
	return
}

// @lc code=end
