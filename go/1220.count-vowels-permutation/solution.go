package countvowelspermutation

/*
 * @lc app=leetcode id=1220 lang=golang
 *
 * [1220] Count Vowels Permutation
 */

// @lc code=start
const MOD int = 1e9 + 7

func countVowelPermutation(n int) (ans int) {
	dp := []int{1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		dpNew := make([]int, 5)
		dpNew[0] = (dp[1] + dp[2] + dp[4]) % MOD
		dpNew[1] = (dp[0] + dp[2]) % MOD
		dpNew[2] = (dp[1] + dp[3]) % MOD
		dpNew[3] = dp[2]
		dpNew[4] = (dp[2] + dp[3]) % MOD

		dp = dpNew
	}

	for _, v := range dp {
		ans = (ans + v) % MOD
	}
	return
}

// @lc code=end
