package minimumflipstomakeaorbequaltoc

/*
 * @lc app=leetcode id=1318 lang=golang
 *
 * [1318] Minimum Flips to Make a OR b Equal to c
 */

// @lc code=start
func minFlips(a int, b int, c int) (ans int) {
	for i := 0; i < 32; i++ {
		ab, bb, cb := a&1, b&1, c&1
		a, b, c = a>>1, b>>1, c>>1

		if cb == 0 {
			if ab == 1 && bb == 1 {
				ans += 2
			} else if ab == 1 || bb == 1 {
				ans += 1
			}
		} else if cb == 1 && ab == 0 && bb == 0 {
			ans += 1
		}
	}
	return
}

// @lc code=end
