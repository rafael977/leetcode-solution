package numberofsubsequencesthatsatisfythegivensumcondition

import (
	"sort"
)

/*
 * @lc app=leetcode id=1498 lang=golang
 *
 * [1498] Number of Subsequences That Satisfy the Given Sum Condition
 */

// @lc code=start
const MOD int = 1e9 + 7

func numSubseq(nums []int, target int) (ans int) {
	sort.Ints(nums)
	f := make([]int, len(nums))
	for i := range f {
		if i == 0 {
			f[i] = 1
		} else {
			f[i] = (f[i-1] * 2) % MOD
		}
	}

	lo, hi := 0, len(nums)-1
	for lo <= hi {
		if nums[lo]+nums[hi] > target {
			hi--
			continue
		}
		ans = (ans + f[hi-lo]) % MOD
		lo++
	}

	return
}

// @lc code=end
