package main

import (
	"fmt"
)

// func fib(n int) int {
// 	if n == 0 || n == 1 {
// 		return n
// 	}

// 	return fib(n-1) + fib(n-2)
// }

func fib(n int) int {
	return helper(n, make(map[int]int))
}

func helper(n int, m map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	var n1, n2 int
	n1, ok := m[n-1]
	if !ok {
		n1 = helper(n-1, m)
	}
	n2, ok = m[n-2]
	if !ok {
		n2 = helper(n-2, m)
	}
	m[n] = n1 + n2
	return n1 + n2
}

func main() {
	fmt.Printf("%v\n", fib(0))
	fmt.Printf("%v\n", fib(1))
	fmt.Printf("%v\n", fib(2))
	fmt.Printf("%v\n", fib(3))
	fmt.Printf("%v\n", fib(4))
	fmt.Printf("%v\n", fib(5))
}
