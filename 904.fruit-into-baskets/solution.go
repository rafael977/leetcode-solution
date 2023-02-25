package fruitintobaskets

/*
 * @lc app=leetcode id=904 lang=golang
 *
 * [904] Fruit Into Baskets
 */

// @lc code=start
func totalFruit(fruits []int) (ans int) {
	cm := make(map[int]int)
	l := 0
	for r, v := range fruits {
		cm[v]++
		for len(cm) > 2 {
			lf := fruits[l]
			cm[lf]--
			if cm[lf] == 0 {
				delete(cm, lf)
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
