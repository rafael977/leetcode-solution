package longesthappystring

/*
 * @lc app=leetcode id=1405 lang=golang
 *
 * [1405] Longest Happy String
 */

// @lc code=start
import (
	pq "github.com/emirpasic/gods/v2/queues/priorityqueue"
)

func longestDiverseString(a int, b int, c int) string {
	type Letter struct {
		l byte
		c int
	}
	q := pq.NewWith(func(x, y *Letter) int {
		if x.c == y.c {
			return int(x.l) - int(y.l)
		}
		return y.c - x.c
	})
	if a > 0 {
		q.Enqueue(&Letter{l: 'a', c: a})
	}
	if b > 0 {
		q.Enqueue(&Letter{l: 'b', c: b})
	}
	if c > 0 {
		q.Enqueue(&Letter{l: 'c', c: c})
	}

	ab := []byte{}
	for !q.Empty() {
		m, _ := q.Dequeue()
		if len(ab) > 0 && m.l == ab[len(ab)-1] {
			s, ok := q.Dequeue()
			if !ok {
				break
			}
			ab = append(ab, s.l)
			s.c--
			if s.c > 0 {
				q.Enqueue(s)
			}
		}
		if m.c >= 2 {
			ab = append(ab, m.l, m.l)
			m.c -= 2
		} else {
			ab = append(ab, m.l)
			m.c -= 1
		}
		if m.c > 0 {
			q.Enqueue(m)
		}
	}

	return string(ab)
}

// @lc code=end
