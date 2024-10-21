package binarytreeswithfactors

import (
	"sort"
)

/*
 * @lc app=leetcode id=823 lang=golang
 *
 * [823] Binary Trees With Factors
 */

// @lc code=start
const MOD int = 1e9 + 7

func numFactoredBinaryTrees(arr []int) (ans int) {

	sort.Ints(arr)
	idx := make(map[int]int)
	for i, v := range arr {
		idx[v] = i
	}
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				if k, ok := idx[arr[i]/arr[j]]; ok {
					dp[i] = (dp[i] + dp[j]*dp[k]) % MOD
				}
			}
		}
	}

	for _, v := range dp {
		ans = (ans + v) % MOD
	}
	return
}

// @lc code=end
