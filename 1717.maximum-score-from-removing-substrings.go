package leetcodesolution

/*
 * @lc app=leetcode id=1717 lang=golang
 *
 * [1717] Maximum Score From Removing Substrings
 */

// @lc code=start
func maximumGain(s string, x int, y int) (ans int) {
	c1, c2 := 'a', 'b'
	p1, p2 := x, y
	if x < y {
		c1, c2 = 'b', 'a'
		p1, p2 = y, x
	}
	st := make([]rune, 0, len(s))
	for _, c := range s {
		if c == c1 || c == c2 {
			if c == c2 && len(st) > 0 && st[len(st)-1] == c1 {
				st = st[:len(st)-1]
				ans += p1
				continue
			}
			st = append(st, c)
		} else {
			c1c, c2c := 0, 0
			for _, v := range st {
				if v == c1 {
					c1c++
				} else if v == c2 {
					c2c++
				}
			}
			ans += p2 * min(c1c, c2c)
			st = st[:0]
		}
	}
	c1c, c2c := 0, 0
	for _, v := range st {
		if v == c1 {
			c1c++
		} else if v == c2 {
			c2c++
		}
	}
	ans += p2 * min(c1c, c2c)
	return
}

// @lc code=end
