package sumofevennumbersafterqueries

/*
 * @lc app=leetcode id=985 lang=golang
 *
 * [985] Sum of Even Numbers After Queries
 */

// @lc code=start
func sumEvenAfterQueries(nums []int, queries [][]int) (ans []int) {
	sum := 0
	for _, v := range nums {
		if v%2 == 0 {
			sum += v
		}
	}

	for _, q := range queries {
		n, np := nums[q[1]], nums[q[1]]+q[0]
		nums[q[1]] = np
		switch {
		case n%2 == 0 && np%2 == 0:
			sum = sum + q[0]
		case n%2 == 0 && np%2 != 0:
			sum = sum - n
		case n%2 != 0 && np%2 == 0:
			sum = sum + np
		}
		ans = append(ans, sum)
	}
	return
}

// @lc code=end
