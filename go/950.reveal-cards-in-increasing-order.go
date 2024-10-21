package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=950 lang=golang
 *
 * [950] Reveal Cards In Increasing Order
 */

// @lc code=start
func deckRevealedIncreasing(deck []int) (ans []int) {
	ans = make([]int, len(deck))
	idx := make([]int, len(deck))
	for i := range idx {
		idx[i] = i
	}
	sort.Ints(deck)
	for _, v := range deck {
		ans[idx[0]] = v
		idx = idx[1:]
		if len(idx) > 0 {
			idx = append(idx, idx[0])
			idx = idx[1:]
		}
	}
	return
}

// @lc code=end
