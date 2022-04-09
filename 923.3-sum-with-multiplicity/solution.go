package sumwithmultiplicity

import "sort"

/*
 * @lc app=leetcode id=923 lang=golang
 *
 * [923] 3Sum With Multiplicity
 */

// @lc code=start
func threeSumMulti(arr []int, target int) int {
	sort.Ints(arr)

	ans := 0
	for i := 0; i < len(arr)-2; i++ {
		r := target - arr[i]
		for j, k := i+1, len(arr)-1; j < k; {
			if arr[j]+arr[k] < r {
				j++
			} else if arr[j]+arr[k] > r {
				k--
			} else {
				if arr[j] == arr[k] {
					ans += (k - j + 1) * (k - j) / 2
					break
				} else {
					cj, ck := 0, 0
					vj, vk := arr[j], arr[k]
					for arr[j] == vj {
						cj++
						j++
					}
					for arr[k] == vk {
						ck++
						k--
					}
					ans += cj * ck
				}
			}
		}
	}

	return ans % 1_000_000_007
}

// @lc code=end
