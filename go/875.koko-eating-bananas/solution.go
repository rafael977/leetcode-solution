package kokoeatingbananas

import "sort"

/*
 * @lc app=leetcode id=875 lang=golang
 *
 * [875] Koko Eating Bananas
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	maxK := 0
	for _, v := range piles {
		if v > maxK {
			maxK = v
		}
	}

	return sort.Search(maxK+1, func(i int) bool {
		if i == 0 {
			return false
		}
		cnt := 0
		for _, v := range piles {
			cnt += v / i
			if v%i != 0 {
				cnt++
			}
		}
		return cnt <= h
	})
}

// @lc code=end
