package sortintegersbythenumberof1bits

import "sort"

/*
 * @lc app=leetcode id=1356 lang=golang
 *
 * [1356] Sort Integers by The Number of 1 Bits
 */

// @lc code=start
func sortByBits(arr []int) []int {
	count := func(v int) int {
		c := 0
		for v > 0 {
			c += v & 1
			v = v >> 1
		}
		return c
	}

	sort.Slice(arr, func(i, j int) bool {
		bi, bj := count(arr[i]), count(arr[j])
		if bi == bj {
			return arr[i] < arr[j]
		}
		return bi < bj
	})
	return arr
}

// @lc code=end
