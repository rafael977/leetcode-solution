package leetcodesolution

/*
 * @lc app=leetcode id=752 lang=golang
 *
 * [752] Open the Lock
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	dm := make(map[string]struct{})
	for _, v := range deadends {
		dm[v] = struct{}{}
	}
	if _, deadlock := dm["0000"]; deadlock {
		return -1
	}

	seen := make(map[string]struct{})
	roll := func(state string) (list []string) {
		bs := []byte(state)
		for i, v := range bs {
			{
				bs[i] = v - 1
				if bs[i] < '0' {
					bs[i] = '9'
				}
				s := string(bs)
				_, didSeen := seen[s]
				_, deadlock := dm[s]
				if !didSeen && !deadlock {
					seen[s] = struct{}{}
					list = append(list, s)
				}
			}
			{
				bs[i] = v + 1
				if bs[i] > '9' {
					bs[i] = '0'
				}
				s := string(bs)
				_, didSeen := seen[s]
				_, deadlock := dm[s]
				if !didSeen && !deadlock {
					seen[s] = struct{}{}
					list = append(list, s)
				}
			}
			bs[i] = v
		}
		return
	}

	q := []string{"0000"}
	seen["0000"] = struct{}{}
	lvl := 0
	for len(q) > 0 {
		nq := make([]string, 0)
		for _, v := range q {
			if v == target {
				return lvl
			}
			nq = append(nq, roll(v)...)
		}
		q = nq
		lvl++
	}
	return -1
}

// @lc code=end
