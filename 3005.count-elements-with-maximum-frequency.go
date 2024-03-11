package leetcodesolution

/*
 * @lc app=leetcode id=3005 lang=golang
 *
 * [3005] Count Elements With Maximum Frequency
 */

// @lc code=start
func maxFrequencyElements(nums []int) (ans int) {
	cm := make(map[int]int)
	for _, v := range nums {
		cm[v]++
	}
	max := 0
	for _, v := range cm {
		if v > max {
			max = v
			ans = 0
		}
		if v == max {
			ans += v
		}
	}
	return
}

// @lc code=end
