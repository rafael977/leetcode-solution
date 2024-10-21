package reducearraysizetothehalf

import "sort"

/*
 * @lc app=leetcode id=1338 lang=golang
 *
 * [1338] Reduce Array Size to The Half
 */

// @lc code=start
func minSetSize(arr []int) int {
	cmap := make(map[int]int)
	for _, v := range arr {
		cmap[v]++
	}

	cnt := make([]int, len(cmap))
	i := 0
	for _, v := range cmap {
		cnt[i] = v
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cnt)))

	rmv, ans := 0, 0
	for _, c := range cnt {
		rmv += c
		ans++
		if rmv >= len(arr)/2 {
			return ans
		}
	}
	return ans
}

// @lc code=end
