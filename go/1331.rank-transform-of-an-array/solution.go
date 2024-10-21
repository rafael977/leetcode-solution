package ranktransformofanarray

import "sort"

/*
 * @lc app=leetcode id=1331 lang=golang
 *
 * [1331] Rank Transform of an Array
 */

// @lc code=start
func arrayRankTransform(arr []int) []int {
	cp := make([]int, len(arr))
	copy(cp, arr)
	sort.Ints(cp)

	m := make(map[int]int)
	i := 1
	for _, v := range cp {
		if _, has := m[v]; !has {
			m[v] = i
			i++
		}
	}

	ans := make([]int, len(arr))
	for i := range arr {
		ans[i] = m[arr[i]]
	}
	return ans
}

// @lc code=end
