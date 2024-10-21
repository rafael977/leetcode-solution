package findthelongestsubstringcontainingvowelsinevencounts

/*
 * @lc app=leetcode id=1371 lang=golang
 *
 * [1371] Find the Longest Substring Containing Vowels in Even Counts
 */

// @lc code=start
var mask = map[byte]int{
	'a': 1,
	'e': 2,
	'i': 4,
	'o': 8,
	'u': 16,
}

func findTheLongestSubstring(s string) (ans int) {
	st := 0
	pos := make([]int, 1<<5)
	for i := range pos {
		pos[i] = -1
	}
	pos[0] = 0
	for i := 0; i < len(s); i++ {
		if v, ok := mask[s[i]]; ok {
			st = st ^ v
		}

		if pos[st] >= 0 {
			ans = max(ans, i+1-pos[st])
		} else {
			pos[st] = i + 1
		}
	}
	return
}

// @lc code=end
