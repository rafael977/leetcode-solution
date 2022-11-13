package containsduplicateii

/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for j, v := range nums {
		if i, ok := m[v]; ok {
			if j-i <= k {
				return true
			}
		}
		m[v] = j
	}
	return false
}

// @lc code=end
