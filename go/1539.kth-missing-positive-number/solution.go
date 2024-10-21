package kthmissingpositivenumber

/*
 * @lc app=leetcode id=1539 lang=golang
 *
 * [1539] Kth Missing Positive Number
 */

// @lc code=start
func findKthPositive(arr []int, k int) int {
	j := 0
	cnt := 0
	for i := 1; i <= k+len(arr); i++ {
		if j < len(arr) && i == arr[j] {
			j++
		} else {
			cnt++
			if cnt == k {
				return i
			}
		}
	}
	return 0
}

// @lc code=end
