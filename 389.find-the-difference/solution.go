package findthedifference

/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	sc, tc := make([]int, 26), make([]int, 26)
	for _, v := range s {
		sc[v-'a']++
	}
	for _, v := range t {
		tc[v-'a']++
	}

	for i := 0; i < 26; i++ {
		if sc[i] != tc[i] {
			return byte(i + 'a')
		}
	}
	return 0
}

// @lc code=end
