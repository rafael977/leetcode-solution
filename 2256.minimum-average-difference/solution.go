package minimumaveragedifference

/*
 * @lc app=leetcode id=2256 lang=golang
 *
 * [2256] Minimum Average Difference
 */

// @lc code=start
func minimumAverageDifference(nums []int) (ans int) {
	mindiff := -1
	n := len(nums)
	f, l := 0, 0
	for _, v := range nums {
		l += v
	}
	for i, v := range nums {
		f += v
		l -= v

		diff := 0
		if i == n-1 {
			diff = absDiff(f/n, 0)
		} else {
			diff = absDiff(f/(i+1), l/(n-i-1))
		}

		if mindiff == -1 || mindiff > diff {
			ans = i
			mindiff = diff
		}
	}
	return
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// @lc code=end
