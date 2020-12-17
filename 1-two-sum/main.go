package main

import "fmt"

func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int, 0)
	ans := make([]int, 2)
	for i, v := range nums {
		if idx, ok := idxMap[v]; ok {
			ans = []int{idx, i}
		}

		idxMap[target-v] = i
	}

	return ans
}

func main() {
	var input []int
	var target int

	input = []int{2, 7, 11, 15}
	target = 9
	fmt.Println(twoSum(input, target))

	input = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(input, target))

	input = []int{3, 3}
	target = 6
	fmt.Println(twoSum(input, target))

	input = []int{1, 2, 3, 4, 5, 6}
	target = 10
	fmt.Println(twoSum(input, target))
}
