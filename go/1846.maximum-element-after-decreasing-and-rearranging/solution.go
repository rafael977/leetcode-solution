package maximumelementafterdecreasingandrearranging

import "sort"

/*
 * @lc app=leetcode id=1846 lang=golang
 *
 * [1846] Maximum Element After Decreasing and Rearranging
 */

// @lc code=start
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	if arr[0] != 1 {
		arr[0] = 1
	}

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}
	return arr[len(arr)-1]
}

// @lc code=end
