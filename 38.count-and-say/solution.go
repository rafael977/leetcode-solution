package countandsay

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	outBuf := strings.Builder{}
	say := countAndSay(n - 1)
	c := say[0]
	cnt := 0
	for _, v := range []byte(say) {
		if c == v {
			cnt++
			continue
		}
		outBuf.WriteString(fmt.Sprintf("%d%c", cnt, c))
		cnt = 1
		c = v
	}
	outBuf.WriteString(fmt.Sprintf("%d%c", cnt, c))
	return outBuf.String()
}

// @lc code=end
