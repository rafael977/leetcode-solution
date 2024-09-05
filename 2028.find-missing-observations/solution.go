package findmissingobservations

/*
 * @lc app=leetcode id=2028 lang=golang
 *
 * [2028] Find Missing Observations
 */

// @lc code=start
func missingRolls(rolls []int, mean int, n int) (ans []int) {
	total := mean * (n + len(rolls))
	mSum := 0
	for _, v := range rolls {
		mSum += v
	}

	nSum := total - mSum
	if nSum > 6*n || nSum < n {
		return
	}
	avg := nSum / n
	r := nSum % n
	ans = make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = avg
		if i < r {
			ans[i]++
		}
	}
	return
}

// @lc code=end
