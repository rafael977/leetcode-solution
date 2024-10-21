package minimummovestoequalarrayelementsii

/*
 * @lc app=leetcode id=462 lang=golang
 *
 * [462] Minimum Moves to Equal Array Elements II
 */

// @lc code=start
func minMoves2(nums []int) (ans int) {
	x := biSearch(nums, 0, len(nums)-1, len(nums)/2)
	for _, v := range nums {
		ans += absDiff(x, v)
	}

	return
}

func biSearch(arr []int, i, j, t int) int {
	if i >= j {
		return arr[i]
	}
	p := partition(arr, i, j)
	if p == t {
		return arr[p]
	} else if p > t {
		return biSearch(arr, i, p-1, t)
	} else {
		return biSearch(arr, p+1, j, t)
	}
}

func partition(arr []int, i, j int) int {
	m := (i + j) / 2
	piv := arr[m]
	arr[m], arr[j] = arr[j], arr[m]

	s := i
	for f := i; f < j; f++ {
		if arr[f] < piv {
			arr[s], arr[f] = arr[f], arr[s]
			s++
		}
	}

	arr[s], arr[j] = arr[j], arr[s]
	return s
}

func absDiff(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

// @lc code=end
