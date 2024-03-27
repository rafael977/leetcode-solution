package leetcodesolution

/*
 * @lc app=leetcode id=442 lang=golang
 *
 * [442] Find All Duplicates in an Array
 */

// @lc code=start
func findDuplicates(nums []int) (ans []int) {
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}
	for k, v := range cnt {
		if v == 2 {
			ans = append(ans, k)
		}
	}
	return
}

// @lc code=end
