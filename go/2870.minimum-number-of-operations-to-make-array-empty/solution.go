package minimumnumberofoperationstomakearrayempty

/*
 * @lc app=leetcode id=2870 lang=golang
 *
 * [2870] Minimum Number of Operations to Make Array Empty
 */

// @lc code=start
func minOperations(nums []int) (ans int) {
	ev := make(map[int]int)
	for _, v := range nums {
		ev[v]++
	}

	for _, v := range ev {
		if v == 1 {
			return -1
		}
		ans += v / 3
		r := v % 3
		if r != 0 {
			ans += 1
		}
	}
	return
}

// @lc code=end
