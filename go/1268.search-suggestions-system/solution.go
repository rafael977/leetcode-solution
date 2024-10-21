package searchsuggestionssystem

import (
	"sort"
)

/*
 * @lc app=leetcode id=1268 lang=golang
 *
 * [1268] Search Suggestions System
 */

// @lc code=start
func suggestedProducts(products []string, searchWord string) (ans [][]string) {
	ans = make([][]string, 0, len(searchWord))
	sort.Strings(products)
	ans = append(ans, products)
	for i := range searchWord {
		ws := make([]string, 0, len(ans[i]))
		for _, w := range ans[i] {
			if i < len(w) && w[i] == searchWord[i] {
				ws = append(ws, w)
			}
		}

		ans = append(ans, ws)
	}

	ans = ans[1:]
	for i := range ans {
		if len(ans[i]) > 3 {
			ans[i] = ans[i][:3]
		}
	}
	return
}

// @lc code=end
