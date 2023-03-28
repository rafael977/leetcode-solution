package canplaceflowers

/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, v := range flowerbed {
		if v == 0 &&
			(i-1 < 0 || (i-1 >= 0 && flowerbed[i-1] == 0)) &&
			(i+1 >= len(flowerbed) || (i+1 < len(flowerbed) && flowerbed[i+1] == 0)) {
			n--
			flowerbed[i] = 1
			if n == 0 {
				return true
			}
		}
	}

	return n <= 0
}

// @lc code=end
