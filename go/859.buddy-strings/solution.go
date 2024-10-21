package buddystrings

/*
 * @lc app=leetcode id=859 lang=golang
 *
 * [859] Buddy Strings
 */

// @lc code=start
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	cm := make(map[byte]int)
	i, j := 0, len(s)-1
	for i < len(s) && s[i] == goal[i] {
		cm[s[i]]++
		i++
	}
	for j >= 0 && s[j] == goal[j] {
		j--
	}

	if i < j {
		sb := []byte(s)
		sb[i], sb[j] = sb[j], sb[i]
		if string(sb) == goal {
			return true
		}
	} else if i > j {
		for _, v := range cm {
			if v > 1 {
				return true
			}
		}
	}
	return false
}

// @lc code=end
