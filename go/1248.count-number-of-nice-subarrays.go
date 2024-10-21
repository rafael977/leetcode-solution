package leetcodesolution

/*
 * @lc app=leetcode id=1248 lang=golang
 *
 * [1248] Count Number of Nice Subarrays
 */

// @lc code=start
func numberOfSubarrays(nums []int, k int) (ans int) {
	sum := 0
	oddSumCount := make(map[int]int)
	oddSumCount[0] = 1 // empty subarray
	for _, v := range nums {
		sum += v & 1
		oddSumCount[sum]++
		if sum >= k {
			ans += oddSumCount[sum-k]
		}
	}

	return
}

// @lc code=end
