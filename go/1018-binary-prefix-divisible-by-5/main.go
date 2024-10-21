package main

import "fmt"

func prefixesDivBy5(A []int) []bool {
	n := 0
	ans := make([]bool, len(A))
	for i, v := range A {
		n = (n<<1 | v) % 5
		if n == 0 {
			ans[i] = true
		}
	}

	return ans
}

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{0, 1, 1, 1, 1, 1}))
}
