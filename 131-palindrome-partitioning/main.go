package main

import (
	"fmt"
)

func partition(s string) [][]string {
	ans := make([][]string, 0)
	backtrack(&ans, s, make([]string, 0))
	return ans
}
func backtrack(ans *[][]string, s string, tmp []string) {
	if len(s) == 0 {
		cp := make([]string, len(tmp))
		copy(cp, tmp)
		*ans = append(*ans, cp)
		return
	}

	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			tmp = append(tmp, s[:i])
			fmt.Println(tmp)
			backtrack(ans, s[i:], tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	// fmt.Printf("%#v", partition("aab"))
	fmt.Printf("%v", partition("cbbbcc"))
}
