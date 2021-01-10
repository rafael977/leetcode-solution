package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	runes := []rune(strings.ToLower(s))
	i, j := 0, len(s)-1

	for i < j {
		if !isAlphanumeric(runes[i]) {
			i++
			continue
		}
		if !isAlphanumeric(runes[j]) {
			j--
			continue
		}

		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("0P"))
}
