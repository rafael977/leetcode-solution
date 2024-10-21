package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=719 lang=golang
 *
 * [719] Find K-th Smallest Pair Distance
 */

// @lc code=start
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0]+1, func(p int) bool {
		cnt, i := 0, 0
		for j, v := range nums {
			for v-nums[i] > p {
				i++
			}
			cnt += j - i
		}

		return cnt >= k
	})
}

// @lc code=end
