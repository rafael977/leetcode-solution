package leetcodesolution

/*
 * @lc app=leetcode id=1404 lang=golang
 *
 * [1404] Number of Steps to Reduce a Number in Binary Representation to One
 */

// @lc code=start
func numSteps(s string) (ans int) {
	sb := []byte(s)
	for len(sb) > 1 {
		if sb[len(sb)-1] == '1' {
			ans++
			i := len(sb) - 1
			for i >= 0 && sb[i] == '1' {
				ans++
				i--
			}
			if i == -1 {
				return
			}
			sb[i] = '1'
			sb = sb[:i+1]
		} else {
			sb = sb[:len(sb)-1]
			ans++
		}
	}
	return
}

// @lc code=end
