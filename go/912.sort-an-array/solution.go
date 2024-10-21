package sortanarray

/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */

// @lc code=start
func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := partition(nums, lo, hi)
	sort(nums, lo, pivot-1)
	sort(nums, pivot+1, hi)
}

func partition(nums []int, lo, hi int) int {
	mid := lo + (hi-lo)/2
	p := nums[mid]
	nums[hi], nums[mid] = nums[mid], nums[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[hi] = nums[hi], nums[i]
	return i
}

// @lc code=end
