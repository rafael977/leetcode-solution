// Created by Rafael Shen at 2024/10/30 10:17
// leetgo: 1.4.9
// https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/

package minimumnumberofremovalstomakemountainarray

// @lc code=begin

func minimumMountainRemovals(nums []int) (ans int) {
	n := len(nums)
	pre := make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				pre[i] = max(pre[i], pre[j]+1)
			}
		}
	}
	suf := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		suf[i] = 1
		for j := n - 1; j > i; j-- {
			if nums[j] < nums[i] {
				suf[i] = max(suf[i], suf[j]+1)
			}
		}
	}

	mx := 0
	for i := 0; i < n; i++ {
		if pre[i] > 1 && suf[i] > 1 {
			mx = max(mx, pre[i]+suf[i]-1)
		}
	}
	ans = n - mx
	return
}

// @lc code=end
