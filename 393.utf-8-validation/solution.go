package utf8validation

/*
 * @lc app=leetcode id=393 lang=golang
 *
 * [393] UTF-8 Validation
 */

// @lc code=start
func validUtf8(data []int) bool {
	i := 0
	for i < len(data) {
		if data[i] <= 127 {
			i++
			continue
		}

		end := i
		switch {
		case data[i] >= 240 && data[i] <= 247:
			end = i + 4
		case data[i] >= 224:
			end = i + 3
		case data[i] >= 192:
			end = i + 2
		default:
			return false
		}
		for i < end {
			if i == len(data) || data[i] < 128 {
				return false
			}
			i++
		}
	}
	return true
}

// @lc code=end
