package main

import (
	"fmt"
	"strings"
)

func LRUCache(strArr []string) string {
	cache := make([]string, 0)
	for i := len(strArr) - 1; i >= 0 && len(cache) < 5; i-- {
		s := strArr[i]
		if !sliceContains(cache, s) {
			cache = append([]string{s}, cache...)
		}
	}
	return strings.Join(cache, "-")

}

func sliceContains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(LRUCache([]string{"A", "B", "A", "C", "A", "B"}))
	fmt.Println(LRUCache([]string{"A", "B", "C", "D", "E", "D", "Q", "Z", "C"}))

}
