package minimumsizesubarraysum

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	ans := len(nums) + 1
	ps := make([]int, len(nums)+1)
	ps[0] = 0
	for i, v := range nums {
		ps[i+1] = ps[i] + v
	}

	for j := 1; j < len(ps); j++ {
		if ps[j] < target {
			continue
		}
		for i := j - 1; i >= 0; i-- {
			if ps[j]-ps[i] >= target {
				if ans > j-i {
					ans = j - i
				}
				break
			}
		}
	}

	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

// @lc code=end
