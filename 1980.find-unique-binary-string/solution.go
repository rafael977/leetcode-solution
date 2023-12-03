package finduniquebinarystring

import "fmt"

/*
 * @lc app=leetcode id=1980 lang=golang
 *
 * [1980] Find Unique Binary String
 */

// @lc code=start
func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	f := fmt.Sprintf("%%0%db", n)
	m := make(map[string]bool)
	for _, v := range nums {
		m[v] = true
	}

	for i := 0; i <= n; i++ {
		bs := fmt.Sprintf(f, i)
		if !m[bs] {
			return bs
		}
	}
	return ""
}

// @lc code=end
