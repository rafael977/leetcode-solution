package maximumproductofwordlengths

/*
 * @lc app=leetcode id=318 lang=golang
 *
 * [318] Maximum Product of Word Lengths
 */

// @lc code=start
func maxProduct(words []string) int {
	dp := make(map[int]uint)

	bLen := func(s string) (r uint) {
		for i := range s {
			r |= 1 << (s[i] - 'a')
		}
		return
	}

	pLen := func(i, j int) int {
		l1, ok := dp[i]
		if !ok {
			l1 = bLen(words[i])
			dp[i] = l1
		}
		l2, ok := dp[j]
		if !ok {
			l2 = bLen(words[j])
			dp[j] = l2
		}

		if l1&l2 == 0 {
			return len(words[i]) * len(words[j])
		}
		return 0
	}

	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			pl := pLen(i, j)
			if max < pl {
				max = pl
			}
		}
	}

	return max
}

// @lc code=end
