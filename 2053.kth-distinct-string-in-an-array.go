package leetcodesolution

/*
 * @lc app=leetcode id=2053 lang=golang
 *
 * [2053] Kth Distinct String in an Array
 */

// @lc code=start
func kthDistinct(arr []string, k int) string {
	sc := make(map[string]int)
	for _, s := range arr {
		sc[s]++
	}

	for _, s := range arr {
		if sc[s] == 1 {
			k--
		}
		if k == 0 {
			return s
		}
	}

	return ""
}

// @lc code=end
