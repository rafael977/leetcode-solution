package permutations

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{append([]int{}, nums...)}
	}

	ans := make([][]int, 0)
	for i, v := range nums {
		next := append([]int{}, nums[:i]...)
		next = append(next, nums[i+1:]...)
		ps := permute(next)
		for _, p := range ps {
			ans = append(ans, append([]int{v}, p...))
		}
	}
	return ans
}

// @lc code=end
