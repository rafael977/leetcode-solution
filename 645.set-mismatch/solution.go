package setmismatch

/*
 * @lc app=leetcode id=645 lang=golang
 *
 * [645] Set Mismatch
 */

// @lc code=start
func findErrorNums(nums []int) (ans []int) {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			ans = append(ans, v)
		}
		m[v] = struct{}{}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			ans = append(ans, i)
		}
	}
	return ans
}

// @lc code=end
