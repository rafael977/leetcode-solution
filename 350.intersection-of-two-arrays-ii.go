package leetcodesolution

/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) (ans []int) {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	for _, v := range nums2 {
		if m[v] > 0 {
			ans = append(ans, v)
			m[v]--
		}
	}

	return
}

// @lc code=end
