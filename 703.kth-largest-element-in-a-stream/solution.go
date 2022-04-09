package kthlargestelementinastream

import (
	"sort"
)

/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 */

// @lc code=start
type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	kl := KthLargest{
		nums: nums,
		k:    k,
	}

	return kl
}

func (kl *KthLargest) Add(val int) int {
	if len(kl.nums) == 0 {
		kl.nums = append(kl.nums, val)
	} else {
		i, j := 0, len(kl.nums)-1
		for i < j {
			m := (i + j) >> 1
			if kl.nums[m] < val {
				i = m + 1
			} else {
				j = m - 1
			}
		}
		if kl.nums[i] < val {
			i++
		}

		kl.nums = append(kl.nums, 0)
		copy(kl.nums[i+1:], kl.nums[i:])
		kl.nums[i] = val
	}

	return kl.nums[len(kl.nums)-kl.k]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
