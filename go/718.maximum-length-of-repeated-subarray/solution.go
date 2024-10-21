package maximumlengthofrepeatedsubarray

/*
 * @lc app=leetcode id=718 lang=golang
 *
 * [718] Maximum Length of Repeated Subarray
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) (ans int) {
	n1, n2 := len(nums1), len(nums2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			}

			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return
}

// @lc code=end
