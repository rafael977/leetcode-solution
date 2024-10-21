package numberoflaserbeamsinabank

/*
 * @lc app=leetcode id=2125 lang=golang
 *
 * [2125] Number of Laser Beams in a Bank
 */

// @lc code=start
func numberOfBeams(bank []string) (ans int) {
	last := 0
	for _, row := range bank {
		c := count(row)
		if c == 0 {
			continue
		}
		if last != 0 {
			ans += last * c
		}
		last = c
	}
	return
}

func count(s string) int {
	c := 0
	for _, v := range s {
		if v == '1' {
			c++
		}
	}
	return c
}

// @lc code=end
