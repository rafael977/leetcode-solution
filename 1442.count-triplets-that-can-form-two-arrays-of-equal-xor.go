package leetcodesolution

/*
 * @lc app=leetcode id=1442 lang=golang
 *
 * [1442] Count Triplets That Can Form Two Arrays of Equal XOR
 */

// @lc code=start
func countTriplets(arr []int) (ans int) {
	for j := 1; j < len(arr); j++ {
		pre := make(map[int]int)
		p := 0
		for i := j - 1; i >= 0; i-- {
			p ^= arr[i]
			pre[p]++
		}

		x := 0
		for k := j; k < len(arr); k++ {
			x ^= arr[k]
			ans += pre[x]
		}
	}
	return
}

// @lc code=end
