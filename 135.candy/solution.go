package candy

/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 */

// @lc code=start
func candy(ratings []int) (ans int) {
	n := len(ratings)
	left := make([]int, n)
	for i := range ratings {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}

		ans += max(right, left[i])
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
