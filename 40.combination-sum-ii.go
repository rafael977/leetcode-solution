package leetcodesolution

import (
	"sort"
)

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	freq := make([][2]int, 0)
	for _, v := range candidates {
		if len(freq) == 0 || freq[len(freq)-1][0] != v {
			freq = append(freq, [2]int{v, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	arr := make([]int, 0)
	var backtrack func(i, t int)
	backtrack = func(i, t int) {
		if t == 0 {
			ans = append(ans, append([]int{}, arr...))
			return
		}
		if i >= len(freq) || freq[i][0] > t {
			return
		}

		backtrack(i+1, t)
		most := min(t/freq[i][0], freq[i][1])
		for j := 1; j <= most; j++ {
			arr = append(arr, freq[i][0])
			backtrack(i+1, t-j*freq[i][0])
		}
		arr = arr[:len(arr)-most]
	}

	backtrack(0, target)
	return
}

// @lc code=end
