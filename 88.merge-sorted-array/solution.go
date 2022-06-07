package mergesortedarray

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	k := len(nums1) - 1
	i, j := m-1, n-1
	for k >= 0 {
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				nums1[k] = nums1[i]
				i--
			} else {
				nums1[k] = nums2[j]
				j--
			}
		} else if j >= 0 {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

// @lc code=end
