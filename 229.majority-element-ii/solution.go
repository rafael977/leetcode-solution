package majorityelementii

/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 */

// @lc code=start
func majorityElement(nums []int) (ans []int) {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	cnt := len(nums) / 3
	for k, v := range m {
		if v > cnt {
			ans = append(ans, k)
		}
	}
	return
}

// @lc code=end
