package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = x * -1
	}

	r := 0
	for x > 0 {
		r = r*10 + x%10
		x = x / 10
	}

	r = r * sign

	if r < math.MinInt32 || r > math.MaxInt32 {
		return 0
	}
	return r
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
}
