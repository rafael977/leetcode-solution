package concatenationofconsecutivebinarynumbers

/*
 * @lc app=leetcode id=1680 lang=golang
 *
 * [1680] Concatenation of Consecutive Binary Numbers
 */

// @lc code=start
const MOD int = 1e9 + 7

func concatenatedBinary(n int) (ans int) {
	shift := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			shift++
		}
		ans = (ans<<shift + i) % MOD
	}
	return
}

// @lc code=end
