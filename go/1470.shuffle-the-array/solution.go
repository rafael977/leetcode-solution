package shufflethearray

/*
 * @lc app=leetcode id=1470 lang=golang
 *
 * [1470] Shuffle the Array
 */

// @lc code=start
func shuffle(nums []int, n int) (ans []int) {
	for i := 0; i < n; i++ {
		ans = append(ans, nums[i])
		ans = append(ans, nums[i+n])
	}
	return
}

// @lc code=end
