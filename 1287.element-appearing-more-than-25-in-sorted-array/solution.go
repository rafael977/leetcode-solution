package elementappearingmorethan25insortedarray

/*
 * @lc app=leetcode id=1287 lang=golang
 *
 * [1287] Element Appearing More Than 25% In Sorted Array
 */

// @lc code=start
func findSpecialInteger(arr []int) int {
	th := len(arr)/4 + 1
	l, r := 0, 0
	for r <= len(arr) {
		if r == len(arr) || arr[l] != arr[r] {
			if r-l >= th {
				return arr[l]
			}
			l = r
		}
		r++
	}
	return 0
}

// @lc code=end
