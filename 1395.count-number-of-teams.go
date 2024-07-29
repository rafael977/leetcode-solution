package leetcodesolution

/*
 * @lc app=leetcode id=1395 lang=golang
 *
 * [1395] Count Number of Teams
 */

// @lc code=start
func numTeams(rating []int) (ans int) {
	n := len(rating)
	for j := 1; j < n; j++ {
		iless, imore := 0, 0
		for i := j - 1; i >= 0; i-- {
			if rating[i] < rating[j] {
				iless++
			}
			if rating[i] > rating[j] {
				imore++
			}
		}
		kless, kmore := 0, 0
		for k := j + 1; k < n; k++ {
			if rating[k] < rating[j] {
				kless++
			}
			if rating[k] > rating[j] {
				kmore++
			}
		}
		ans += iless*kmore + imore*kless
	}
	return
}

// @lc code=end
