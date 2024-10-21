package main

import "fmt"

type input struct {
	haystack string
	needle   string
}

type test struct {
	input
	output int
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		var j int
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}

		if len(needle) == j {
			return i
		}
	}
	return -1
}

func main() {
	tests := []test{
		{input{"hello", "ll"}, 2},
		{input{"aaaaa", "bba"}, -1},
		{input{"aaaaa", "bba"}, -1},
	}

	for _, v := range tests {
		if strStr(v.haystack, v.needle) != v.output {
			fmt.Println("Test failed", v)
		}
	}
	fmt.Println("Test successful")
}
