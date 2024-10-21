package ransomnote

/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	rm, mm := make([]int, 26), make([]int, 26)
	for _, c := range ransomNote {
		rm[c-'a']++
	}
	for _, c := range magazine {
		mm[c-'a']++
	}

	for i := range rm {
		if rm[i] > mm[i] {
			return false
		}
	}
	return true
}

// @lc code=end
