package largestnumber

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

// @lc code=start
func largestNumber(nums []int) (ans string) {
	sort.Slice(nums, func(i, j int) bool {
		return fmt.Sprintf("%d%d", nums[i], nums[j]) > fmt.Sprintf("%d%d", nums[j], nums[i])
	})

	numsText := make([]string, len(nums))
	for i, v := range nums {
		numsText[i] = strconv.Itoa(v)
	}

	ans = strings.TrimLeft(strings.Join(numsText, ""), "0")
	if ans == "" {
		return "0"
	}
	return
}

// @lc code=end
