package numberswithsameconsecutivedifferences

/*
 * @lc app=leetcode id=967 lang=golang
 *
 * [967] Numbers With Same Consecutive Differences
 */

// @lc code=start
func numsSameConsecDiff(n int, k int) []int {
	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n--
	for n > 0 {
		newq := make([]int, 0)
		for _, v := range q {
			dig := v % 10
			if dig-k >= 0 && dig-k <= 9 {
				newq = append(newq, v*10+dig-k)
			}
			if k != 0 && dig+k >= 0 && dig+k <= 9 {
				newq = append(newq, v*10+dig+k)
			}
		}
		q = newq
		n--
	}
	return q
}

// @lc code=end
