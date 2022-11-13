package groupanagrams

import "sort"

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) (ans [][]string) {
	m := make(map[string][]string)
	for _, v := range strs {
		dv := []byte(v)
		sort.Slice(dv, func(i, j int) bool {
			return dv[i] < dv[j]
		})
		dvs := string(dv)
		if _, ok := m[dvs]; !ok {
			m[dvs] = make([]string, 0)
		}
		m[dvs] = append(m[dvs], v)
	}
	for _, v := range m {
		ans = append(ans, v)
	}
	return
}

// @lc code=end
