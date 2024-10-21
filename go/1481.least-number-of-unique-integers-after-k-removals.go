package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=1481 lang=golang
 *
 * [1481] Least Number of Unique Integers after K Removals
 */

// @lc code=start
func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	uv := make([]int, 0, len(m))
	for _, v := range m {
		uv = append(uv, v)
	}

	sort.Ints(uv)
	i := 0
	for k > 0 && i < len(uv) {
		if k-uv[i] < 0 {
			break
		}
		k -= uv[i]
		i++
	}
	return len(uv) - i
}

// @lc code=end
