package stringcompressionq

import "fmt"

/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */

// @lc code=start
func compress(chars []byte) int {
	com := ""
	c := chars[0]
	cnt := 1

	output := func(c byte, cnt int) string {
		if cnt == 1 {
			return string(c)
		}
		return fmt.Sprintf("%c%d", c, cnt)
	}

	for _, v := range chars[1:] {
		if c != v {
			com = com + output(c, cnt)
			c = v
			cnt = 1
		} else {
			cnt++
		}
	}
	com = com + output(c, cnt)
	for i := 0; i < len(com); i++ {
		chars[i] = com[i]
	}
	return len(com)
}

// @lc code=end
