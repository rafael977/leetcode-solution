package xorqueriesofasubarray

/*
 * @lc app=leetcode id=1310 lang=golang
 *
 * [1310] XOR Queries of a Subarray
 */

// @lc code=start
func xorQueries(arr []int, queries [][]int) []int {
	pxor := make([]int, len(arr))
	pxor[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		pxor[i] = pxor[i-1] ^ arr[i]
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		if q[0] == 0 {
			ans[i] = pxor[q[1]]
		} else {
			ans[i] = pxor[q[1]] ^ pxor[q[0]-1]
		}
	}

	return ans
}

// @lc code=end
