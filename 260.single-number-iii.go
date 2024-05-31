package leetcodesolution

/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 */

// @lc code=start
func singleNumber(nums []int) (ans []int) {
	xorAll := 0
	for _, v := range nums {
		xorAll ^= v
	}

	lowbit := xorAll & -xorAll
	ans = make([]int, 2)
	for _, v := range nums {
		if v&lowbit == 0 {
			ans[0] ^= v
		} else {
			ans[1] ^= v
		}
	}
	return
}

// @lc code=end
