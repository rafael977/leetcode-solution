package findfirstandlastpositionofelementinsortedarray

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) (ans []int) {
	ans = []int{-1, -1}

	n := len(nums)
	if n == 0 {
		return
	}

	for i, j := 0, n-1; i <= j; {
		p := (i + j) / 2
		if nums[p] == target && (p == 0 || nums[p-1] != target) {
			ans[0] = p
			break
		}
		if nums[p] >= target {
			j = p - 1
		} else {
			i = p + 1
		}
	}
	for i, j := 0, n-1; i <= j; {
		p := (i + j) / 2
		if nums[p] == target && (p == n-1 || nums[p+1] != target) {
			ans[1] = p
			break
		}
		if nums[p] > target {
			j = p - 1
		} else {
			i = p + 1
		}
	}
	return
}

// @lc code=end
