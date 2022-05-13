package permutationsii

import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)

	n := len(nums)
	visited := make([]bool, n)
	perm := make([]int, 0)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int{}, perm...))
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] || i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}

			visited[i] = true
			perm = append(perm, nums[i])

			backtrack(idx + 1)

			visited[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

// @lc code=end
