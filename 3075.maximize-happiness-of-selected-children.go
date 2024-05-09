package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=3075 lang=golang
 *
 * [3075] Maximize Happiness of Selected Children
 */

// @lc code=start
func maximumHappinessSum(happiness []int, k int) (ans int64) {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	for i := 0; i < k; i++ {
		if happiness[i]-i <= 0 {
			break
		}
		ans += int64(happiness[i] - i)
	}
	return
}

// @lc code=end
