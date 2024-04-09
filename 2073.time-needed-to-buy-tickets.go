package leetcodesolution

/*
 * @lc app=leetcode id=2073 lang=golang
 *
 * [2073] Time Needed to Buy Tickets
 */

// @lc code=start
func timeRequiredToBuy(tickets []int, k int) int {
	n := len(tickets)
	t := n * tickets[k]

	for i, v := range tickets {
		if i == k {
			continue
		}
		if v < tickets[k] {
			t -= tickets[k] - v
		} else if i > k {
			t -= 1
		}
	}
	return t
}

// @lc code=end
