package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, v := range wordDict {
		m[v] = true
	}

	return helper(s, m)
}
func helper(s string, m map[string]bool) bool {
	if len(s) == 0 {
		return false
	}
	if m[s] {
		return true
	}

	for i := 0; i < len(s); i++ {
		if m[s[:i]] && helper(s[i:], m) {
			return true
		}
	}

	return false
}

func main() {
	// fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	// fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	// fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}
