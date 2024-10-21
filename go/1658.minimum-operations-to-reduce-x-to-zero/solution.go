package minimumoperationstoreducextozero

/*
 * @lc app=leetcode id=1658 lang=golang
 *
 * [1658] Minimum Operations to Reduce X to Zero
 */

// @lc code=start
func minOperations(nums []int, x int) int {
	numsSum := 0
	for _, v := range nums {
		numsSum += v
	}

	if numsSum < x {
		return -1
	}

	targetSum := numsSum - x
	l, r := 0, 0
	sum := 0
	ml := -1

	for r < len(nums) {
		sum += nums[r]
		for sum > targetSum {
			sum -= nums[l]
			l++
		}

		if sum == targetSum {
			ml = max(ml, r-l+1)
		}

		r++
	}

	if ml == -1 {
		return -1
	}
	return len(nums) - ml
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
