package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		b := math.Sqrt(float64(c - a*a))
		if b == math.Floor(b) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(4))
	fmt.Println(judgeSquareSum(3))
	fmt.Println(judgeSquareSum(2))
	fmt.Println(judgeSquareSum(1))
}
