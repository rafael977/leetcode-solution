package combinationsumiii

/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start
func combinationSum3(k int, n int) (ans [][]int) {
	ans = make([][]int, 0)
	var backtrack func(cur []int, j, k, n int)
	backtrack = func(cur []int, j, k, n int) {
		if j > 9 || k == 0 || n < 1 {
			return
		}
		if k == 1 && n == j {
			comb := append([]int{}, cur...)
			ans = append(ans, append(comb, j))
			return
		}

		cur = append(cur, j)
		backtrack(cur, j+1, k-1, n-j)
		cur = cur[:len(cur)-1]
		backtrack(cur, j+1, k, n)
	}

	backtrack([]int{}, 1, k, n)
	return
}

// @lc code=end
