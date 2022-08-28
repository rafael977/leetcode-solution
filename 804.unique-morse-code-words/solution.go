package uniquemorsecodewords

import "strings"

/*
 * @lc app=leetcode id=804 lang=golang
 *
 * [804] Unique Morse Code Words
 */

// @lc code=start
var morse []string = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	tmap := make(map[string]bool)
	for _, w := range words {
		t := strings.Builder{}
		for _, c := range w {
			t.WriteString(morse[c-'a'])
		}
		tmap[t.String()] = true
	}

	return len(tmap)
}

// @lc code=end
