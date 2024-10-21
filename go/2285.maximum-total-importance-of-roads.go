package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=2285 lang=golang
 *
 * [2285] Maximum Total Importance of Roads
 */

// @lc code=start
func maximumImportance(n int, roads [][]int) int64 {
	cc := make([]int, n)
	for _, r := range roads {
		cc[r[0]]++
		cc[r[1]]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cc)))

	ans := 0
	for i := range cc {
		ans += cc[i] * (n - i)
	}

	return int64(ans)
}

// @lc code=end
