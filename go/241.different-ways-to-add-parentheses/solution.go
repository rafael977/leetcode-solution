package differentwaystoaddparentheses

/*
 * @lc app=leetcode id=241 lang=golang
 *
 * [241] Different Ways to Add Parentheses
 */

// @lc code=start
func diffWaysToCompute(expression string) []int {
	var dfs func(l, r int) []int
	dfs = func(l, r int) []int {
		ans := make([]int, 0)
		for i := l; i <= r; i++ {
			if expression[i] >= '0' && expression[i] <= '9' {
				continue
			}
			l1, l2 := dfs(l, i-1), dfs(i+1, r)
			for _, v1 := range l1 {
				for _, v2 := range l2 {
					if expression[i] == '+' {
						ans = append(ans, v1+v2)
					} else if expression[i] == '-' {
						ans = append(ans, v1-v2)
					} else {
						ans = append(ans, v1*v2)
					}
				}
			}
		}
		if len(ans) == 0 {
			n := 0
			for i := l; i <= r; i++ {
				n = n*10 + int(expression[i]-'0')
			}
			ans = append(ans, n)
		}
		return ans
	}

	return dfs(0, len(expression)-1)
}

// @lc code=end
