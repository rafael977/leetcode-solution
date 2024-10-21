package dividearrayintoarrayswithmaxdifference

import "sort"

/*
 * @lc app=leetcode id=2966 lang=golang
 *
 * [2966] Divide Array Into Arrays With Max Difference
 */

// @lc code=start
func divideArray(nums []int, k int) (ans [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		ans = append(ans, append([]int{}, nums[i:i+3]...))
	}
	return
}

// @lc code=end
