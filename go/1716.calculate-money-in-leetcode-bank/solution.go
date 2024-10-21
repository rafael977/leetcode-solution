package calculatemoneyinleetcodebank

/*
 * @lc app=leetcode id=1716 lang=golang
 *
 * [1716] Calculate Money in Leetcode Bank
 */

// @lc code=start
const aWeek = 1 + 2 + 3 + 4 + 5 + 6 + 7

func totalMoney(n int) (ans int) {
	weeks, days := n/7, n%7
	ans += aWeek*weeks + (weeks-1)*weeks/2*7
	ans += weeks*days + (1+days)*days/2
	return
}

// @lc code=end
