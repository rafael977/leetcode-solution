// Created by Rafael Shen at 2024/10/28 10:10
// leetgo: 1.4.9
// https://leetcode.com/problems/longest-square-streak-in-an-array/

package longestsquarestreakinanarray

import "sort"

// @lc code=begin

func longestSquareStreak(nums []int) (ans int) {
	sort.Ints(nums)
	m := make(map[int]int)
	for _, v := range nums {
		x := m[v] + 1
		ans = max(ans, x)
		m[v*v] = x
	}
	if ans < 2 {
		ans = -1
	}
	return
}

// @lc code=end
