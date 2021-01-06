package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res ^ v
	}

	return res
}

func main() {
	nums := []int{2, 2, 1}
	fmt.Printf("%v\n", singleNumber(nums))

	nums = []int{4, 1, 2, 2, 1}
	fmt.Printf("%v\n", singleNumber(nums))

	nums = []int{1}
	fmt.Printf("%v\n", singleNumber(nums))
}
