package main

import "fmt"

func decode(encoded []int, first int) []int {
	nums := make([]int, len(encoded)+1)
	nums[0] = first

	for i, v := range encoded {
		nums[i+1] = nums[i] ^ v
	}

	return nums
}

func main() {
	fmt.Println(decode([]int{1, 2, 3}, 1))
	fmt.Println(decode([]int{6, 2, 7, 3}, 4))
}
