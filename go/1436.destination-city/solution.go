package destinationcity

/*
 * @lc app=leetcode id=1436 lang=golang
 *
 * [1436] Destination City
 */

// @lc code=start
func destCity(paths [][]string) string {
	m := make(map[string]int)
	for _, p := range paths {
		m[p[0]]++
		m[p[1]]--
	}

	for k, v := range m {
		if v < 0 {
			return k
		}
	}
	return ""
}

// @lc code=end
