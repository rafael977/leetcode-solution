package countnumberofmaximumbitwiseorsubsets

/*
 * @lc app=leetcode id=2044 lang=golang
 *
 * [2044] Count Number of Maximum Bitwise-OR Subsets
 */

// @lc code=start
func countMaxOrSubsets(nums []int) (ans int) {
	n := len(nums)
	mor := 0
	for _, v := range nums {
		mor |= v
	}
	for i := 0; i < 1<<n; i++ {
		or := 0
		mask := 1
		for j := 0; j < n; j++ {
			if i&(mask<<j) != 0 {
				or |= nums[j]
			}
		}
		if mor == or {
			ans++
		}
	}

	return
}

// @lc code=end
