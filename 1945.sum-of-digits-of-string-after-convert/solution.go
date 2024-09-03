package sumofdigitsofstringafterconvert

/*
 * @lc app=leetcode id=1945 lang=golang
 *
 * [1945] Sum of Digits of String After Convert
 */

// @lc code=start
func getLucky(s string, k int) (ans int) {
	for _, v := range s {
		n := int(v - 'a' + 1)
		ans += n/10 + n%10
	}
	for i := 1; i < k; i++ {
		old := ans
		ans = 0
		for old > 0 {
			ans += old % 10
			old = old / 10
		}
	}
	return
}

// @lc code=end
