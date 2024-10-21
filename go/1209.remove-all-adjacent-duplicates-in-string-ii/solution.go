package removealladjacentduplicatesinstringii

/*
 * @lc app=leetcode id=1209 lang=golang
 *
 * [1209] Remove All Adjacent Duplicates in String II
 */

// @lc code=start
func removeDuplicates(s string, k int) string {
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] != s[i-1] {
			stack = append(stack, 1)
		} else {
			if stack[len(stack)-1]+1 == k {
				stack = stack[:len(stack)-1]
				s = s[:i-k+1] + s[i+1:]
				i = i - k
			} else {
				stack[len(stack)-1]++
			}
		}
	}

	return s
}

// @lc code=end
