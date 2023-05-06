package findthedifferenceoftwoarrays

/*
 * @lc app=leetcode id=2215 lang=golang
 *
 * [2215] Find the Difference of Two Arrays
 */

// @lc code=start
func findDifference(nums1 []int, nums2 []int) (ans [][]int) {
	ans = make([][]int, 2)
	ans[0] = make([]int, 0)
	ans[1] = make([]int, 0)
	m1, m2 := make(map[int]bool), make(map[int]bool)
	for _, v := range nums1 {
		m1[v] = true
	}
	for _, v := range nums2 {
		m2[v] = true
	}
	for k := range m1 {
		if !m2[k] {
			ans[0] = append(ans[0], k)
		}
	}
	for k := range m2 {
		if !m1[k] {
			ans[1] = append(ans[1], k)
		}
	}
	return
}

// @lc code=end
