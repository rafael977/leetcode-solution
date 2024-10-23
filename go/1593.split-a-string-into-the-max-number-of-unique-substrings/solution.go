// Created by Rafael Shen at 2024/10/21 12:54
// leetgo: 1.4.9
// https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/

package splitastringintothemaxnumberofuniquesubstrings

// @lc code=begin

func maxUniqueSplit(s string) (ans int) {
	m := make(map[string]struct{})
	var backtrack func(idx, n int)
	backtrack = func(idx, n int) {
		if idx >= len(s) {
			ans = max(ans, n)
		}

		for i := idx; i < len(s); i++ {
			if _, has := m[s[idx:i+1]]; !has {
				m[s[idx:i+1]] = struct{}{}
				backtrack(i+1, n+1)
				delete(m, s[idx:i+1])
			}
		}
	}

	backtrack(0, 0)
	return
}

// @lc code=end
