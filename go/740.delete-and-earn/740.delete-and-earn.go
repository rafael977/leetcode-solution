package main

/*
 * @lc app=leetcode id=740 lang=golang
 *
 * [740] Delete and Earn
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxN := 0
	for _, n := range nums {
		if n > maxN {
			maxN = n
		}
	}
	sums := make([]int, maxN+1)
	for _, n := range nums {
		sums[n] += n
	}

	rob := func(sums []int) int {
		first, second := sums[0], max(sums[0], sums[1])
		for i := 2; i <= maxN; i++ {
			first, second = second, max(sums[i]+first, second)
		}
		return second
	}

	return rob(sums)
}

// @lc code=end
