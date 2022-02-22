package main

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}

	var bfs func(tar, idx int)
	bfs = func(tar, idx int) {
		if idx >= len(candidates) {
			return
		}
		if tar == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		bfs(tar, idx+1)
		if tar-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			bfs(tar-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}

	bfs(target, 0)
	return
}

// @lc code=end
