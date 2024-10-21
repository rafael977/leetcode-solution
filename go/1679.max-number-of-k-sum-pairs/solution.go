package maxnumberofksumpairs

/*
 * @lc app=leetcode id=1679 lang=golang
 *
 * [1679] Max Number of K-Sum Pairs
 */

// @lc code=start
func maxOperations(nums []int, k int) (ans int) {
	m := make(map[int]int)
	for _, v := range nums {

		if m[v] > 0 {
			ans++
			m[v]--
		} else {
			m[k-v]++
		}
	}
	return
}

// @lc code=end
