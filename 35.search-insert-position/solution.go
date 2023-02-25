package searchinsertposition

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) (ans int) {
	lo, hi := 0, len(nums)-1
	ans = len(nums)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if target <= nums[mid] {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return
}

// @lc code=end
