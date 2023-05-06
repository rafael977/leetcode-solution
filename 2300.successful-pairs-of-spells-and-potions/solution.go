package successfulpairsofspellsandpotions

import "sort"

/*
 * @lc app=leetcode id=2300 lang=golang
 *
 * [2300] Successful Pairs of Spells and Potions
 */

// @lc code=start
func successfulPairs(spells []int, potions []int, success int64) []int {
	for i := range potions {
		potions[i] = (int(success) + potions[i] - 1) / potions[i]
	}
	sort.Ints(potions)
	ans := make([]int, len(spells))
	for i := range ans {
		idx := sort.Search(len(potions), func(j int) bool {
			return spells[i] < potions[j]
		})
		if idx == -1 {
			ans[i] = len(potions)
		} else {
			ans[i] = idx
		}
	}
	return ans
}

// @lc code=end
