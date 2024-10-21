package leetcodesolution

/*
 * @lc app=leetcode id=2149 lang=golang
 *
 * [2149] Rearrange Array Elements by Sign
 */

// @lc code=start
func rearrangeArray(nums []int) (ans []int) {
	i, j := 0, 1
	ans = make([]int, len(nums))
	for _, v := range nums {
		if v > 0 {
			ans[i] = v
			i += 2
		} else {
			ans[j] = v
			j += 2
		}
	}
	return
}

// @lc code=end
