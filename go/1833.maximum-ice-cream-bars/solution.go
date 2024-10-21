package maximumicecreambars

import "sort"

/*
 * @lc app=leetcode id=1833 lang=golang
 *
 * [1833] Maximum Ice Cream Bars
 */

// @lc code=start
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	sum := 0
	for i, v := range costs {
		sum += v
		if sum > coins {
			return i
		}
	}
	return len(costs)
}

// @lc code=end
