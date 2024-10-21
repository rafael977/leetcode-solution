package parsingabooleanexpression

/*
 * @lc app=leetcode id=1106 lang=golang
 *
 * [1106] Parsing A Boolean Expression
 */

// @lc code=start
func parseBoolExpr(expression string) bool {
	and := func(tf []byte) byte {
		for _, v := range tf {
			if v == 'f' {
				return 'f'
			}
		}
		return 't'
	}
	or := func(tf []byte) byte {
		for _, v := range tf {
			if v == 't' {
				return 't'
			}
		}
		return 'f'
	}
	not := func(tf byte) byte {
		if tf == 't' {
			return 'f'
		}
		return 't'
	}
	st := make([]byte, 0, len(expression))
	for _, v := range []byte(expression) {
		if v == '&' || v == '|' || v == '!' ||
			v == 't' || v == 'f' {
			st = append(st, v)
		} else if v == ')' {
			tf := make([]byte, 0)
			for st[len(st)-1] == 't' || st[len(st)-1] == 'f' {
				tf = append(tf, st[len(st)-1])
				st = st[:len(st)-1]
			}
			op := st[len(st)-1]
			st = st[:len(st)-1]
			if op == '&' {
				st = append(st, and(tf))
			} else if op == '|' {
				st = append(st, or(tf))
			} else {
				st = append(st, not(tf[0]))
			}
		}
	}

	if st[0] == 't' {
		return true
	}
	return false
}

// @lc code=end
