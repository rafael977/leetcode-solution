package maximumpointsyoucanobtainfromcards

/*
 * @lc app=leetcode id=1423 lang=golang
 *
 * [1423] Maximum Points You Can Obtain from Cards
 */

// @lc code=start
func maxScore(cardPoints []int, k int) int {
	total := 0
	for _, v := range cardPoints {
		total += v
	}

	l := len(cardPoints) - k
	sum := 0
	for i := 0; i < l; i++ {
		sum += cardPoints[i]
	}
	min := sum
	for i := l; i < len(cardPoints); i++ {
		sum += cardPoints[i] - cardPoints[i-l]
		if sum < min {
			min = sum
		}
	}

	return total - min
}

// @lc code=end
