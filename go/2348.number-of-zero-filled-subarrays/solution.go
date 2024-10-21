package numberofzerofilledsubarrays

/*
 * @lc app=leetcode id=2348 lang=golang
 *
 * [2348] Number of Zero-Filled Subarrays
 */

// @lc code=start
func zeroFilledSubarray(nums []int) int64 {
	zc := 0
	ans := 0
	for _, v := range nums {
		if v != 0 {
			if zc > 0 {
				ans += count(zc)
				zc = 0
			}
		} else {
			zc++
		}
	}
	if zc > 0 {
		ans += count(zc)
	}
	return int64(ans)
}

func count(n int) int {
	return (n + 1) * n / 2
}

// @lc code=end
