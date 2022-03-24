package main

import "strings"

/*
 * @lc app=leetcode id=1663 lang=golang
 *
 * [1663] Smallest String With A Given Numeric Value
 */

// @lc code=start
func getSmallestString(n int, k int) string {
	zs := 0
	for k >= n-1+26 {
		zs++
		k -= 26
		n--
	}

	var sb strings.Builder
	if n > 0 {
		for i := 0; i < n-1; i++ {
			sb.WriteByte('a')
		}
		sb.WriteByte(byte('a' + k - n))
	}

	for i := 0; i < zs; i++ {
		sb.WriteByte('z')
	}
	return sb.String()
}

// @lc code=end
