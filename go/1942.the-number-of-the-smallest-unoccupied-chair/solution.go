package thenumberofthesmallestunoccupiedchair

/*
 * @lc app=leetcode id=1942 lang=golang
 *
 * [1942] The Number of the Smallest Unoccupied Chair
 */

// @lc code=start

import (
	"sort"

	pq "github.com/emirpasic/gods/v2/queues/priorityqueue"
)

func smallestChair(times [][]int, targetFriend int) int {
	type Friend struct {
		idx, arrival, leaving, seat int
	}
	fs := make([]*Friend, 0, len(times))
	for i, t := range times {
		fs = append(fs, &Friend{
			idx:     i,
			arrival: t[0],
			leaving: t[1],
		})
	}
	sort.Slice(fs, func(i, j int) bool {
		return fs[i].arrival < fs[j].arrival
	})

	fq := pq.NewWith(func(x, y *Friend) int {
		return x.leaving - y.leaving
	})
	sq := pq.New[int]()
	s := 0
	for _, v := range fs {
		for !fq.Empty() {
			if h, _ := fq.Peek(); h.leaving <= v.arrival {
				sq.Enqueue(h.seat)
				fq.Dequeue()
			} else {
				break
			}
		}
		if !sq.Empty() {
			sn, _ := sq.Dequeue()
			v.seat = sn
		} else {
			v.seat = s
			s++
		}

		if v.idx == targetFriend {
			return v.seat
		}

		fq.Enqueue(v)
	}
	return -1
}

// @lc code=end
