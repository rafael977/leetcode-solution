package numberofdicerollswithtargetsum

/*
 * @lc app=leetcode id=1155 lang=golang
 *
 * [1155] Number of Dice Rolls With Target Sum
 */

// @lc code=start
const MOD int = 1e9 + 7

func numRollsToTarget(n int, k int, target int) int {
	r := map[int]int{target: 1}
	for i := 0; i < n; i++ {
		nr := make(map[int]int)
		for x, c := range r {
			for j := 1; j <= k; j++ {
				if x-j >= 0 {
					nr[x-j] = (nr[x-j] + c) % MOD
				}
			}
		}
		r = nr
	}
	return r[0]
}

// @lc code=end
