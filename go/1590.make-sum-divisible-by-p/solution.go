package makesumdivisiblebyp

/*
 * @lc app=leetcode id=1590 lang=golang
 *
 * [1590] Make Sum Divisible by P
 */

// @lc code=start
func minSubarray(nums []int, p int) int {
	psum := make([]int, len(nums)+1)
	for i, v := range nums {
		psum[i+1] = psum[i] + v
	}
	r := psum[len(psum)-1] % p
	if r == 0 {
		return 0
	}

	for l := 1; l < len(nums); l++ {
		for i := 0; i < len(psum)-l; i++ {
			if (psum[i+l]-psum[i])%p == r {
				return l
			}
		}
	}
	return -1
}

// @lc code=end
