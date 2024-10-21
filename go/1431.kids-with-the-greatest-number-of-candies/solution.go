package kidswiththegreatestnumberofcandies

/*
 * @lc app=leetcode id=1431 lang=golang
 *
 * [1431] Kids With the Greatest Number of Candies
 */

// @lc code=start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for _, v := range candies[1:] {
		if v > max {
			max = v
		}
	}

	ans := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= max {
			ans[i] = true
		}
	}
	return ans
}

// @lc code=end
