package checkifastringcontainsallbinarycodesofsizek

/*
 * @lc app=leetcode id=1461 lang=golang
 *
 * [1461] Check If a String Contains All Binary Codes of Size K
 */

// @lc code=start
func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}

	n := 0
	for i := 0; i <= k-1; i++ {
		n <<= 1
		if s[i] == '1' {
			n++
		}
	}

	mask := (1 << (k - 1)) - 1
	cnt := make(map[int]bool)
	cnt[n] = true
	for i := 1; i <= len(s)-k; i++ {
		n = (n & mask) << 1
		if s[i+k-1] == '1' {
			n++
		}
		cnt[n] = true
	}

	return len(cnt) == 1<<k
}

// @lc code=end
