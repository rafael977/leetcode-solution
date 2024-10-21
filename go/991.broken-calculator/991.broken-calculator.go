package main

/*
 * @lc app=leetcode id=991 lang=golang
 *
 * [991] Broken Calculator
 */

// @lc code=start
func brokenCalc(startValue int, target int) (ans int) {
	for startValue < target {
		if target%2 == 1 {
			target++
		} else {
			target /= 2
		}
		ans++
		// if (startValue-1)*2 >= target {
		// 	startValue--
		// } else {
		// 	startValue *= 2
		// }
		// ans++
	}
	ans += startValue - target
	return
}

// @lc code=end
