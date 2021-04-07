package main

func search(nums []int, target int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}

	return false
}

func main() {
	println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}
