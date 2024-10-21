package dailytemperatures

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	s := make([]int, 0)
	for j := range temperatures {
		for len(s) > 0 {
			i := s[len(s)-1]
			if temperatures[i] >= temperatures[j] {
				break
			}
			ans[i] = j - i
			s = s[:len(s)-1]
		}
		s = append(s, j)
	}
	return ans
}

// @lc code=end
