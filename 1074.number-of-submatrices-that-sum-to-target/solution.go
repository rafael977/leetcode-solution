package numberofsubmatricesthatsumtotarget

/*
 * @lc app=leetcode id=1074 lang=golang
 *
 * [1074] Number of Submatrices That Sum to Target
 */

// @lc code=start
func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
	for i := range matrix {
		sums := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] { // column sums of rows from i
			for c, v := range row {
				sums[c] += v
			}

			ans += subarraySum(sums, target)
		}
	}
	return
}

func subarraySum(nums []int, k int) (ans int) {
	preSum := map[int]int{0: 1}
	sum := 0
	for _, n := range nums {
		sum += n
		if _, ok := preSum[sum-k]; ok {
			ans += preSum[sum-k]
		}
		preSum[sum]++
	}
	return
}

// @lc code=end
