package addtoarrayformofinteger

/*
 * @lc app=leetcode id=989 lang=golang
 *
 * [989] Add to Array-Form of Integer
 */

// @lc code=start
func addToArrayForm(num []int, k int) (ans []int) {
	n := len(num)
	carry := 0
	for i := 0; i < n || k > 0; i++ {
		if i < n {
			carry += num[n-1-i]
		}
		if k > 0 {
			carry += (k % 10)
			k /= 10
		}
		ans = append([]int{carry % 10}, ans...)
		carry /= 10
	}
	if carry > 0 {
		ans = append([]int{carry % 10}, ans...)
	}

	return
}

// @lc code=end
