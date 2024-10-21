package leetcodesolution

import "math"

/*
 * @lc app=leetcode id=2976 lang=golang
 *
 * [2976] Minimum Cost to Convert String I
 */

// @lc code=start
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) (ans int64) {
	mp := make([][]int, 26)
	for i := range mp {
		mp[i] = make([]int, 26)
		for j := range mp[i] {
			mp[i][j] = math.MaxInt32 / 2
		}
	}
	for i, c := range cost {
		mp[original[i]-'a'][changed[i]-'a'] = min(mp[original[i]-'a'][changed[i]-'a'], c)
	}

	for k := 0; k < 26; k++ {
		mp[k][k] = 0
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				mp[i][j] = min(mp[i][j], mp[i][k]+mp[k][j])
			}
		}
	}

	for i := range source {
		if source[i] == target[i] {
			continue
		}

		if mp[source[i]-'a'][target[i]-'a'] == math.MaxInt32/2 {
			return -1
		}
		ans += int64(mp[source[i]-'a'][target[i]-'a'])
	}
	return
}

// @lc code=end
