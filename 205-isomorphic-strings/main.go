package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	l := len(s)

	sm, tm := make(map[byte]int, 0), make(map[byte]int, 0)
	for i := 0; i < l; i++ {
		sc, oks := sm[s[i]]
		tc, okt := tm[t[i]]

		if !oks && !okt {
			sm[s[i]] = i
			tm[t[i]] = i
			continue
		} else if oks && okt {
			if sc == tc {
				continue
			} else {
				return false
			}

		} else {
			return false
		}
	}

	return true
}

func main() {
	var s, t string

	s, t = "egg", "add"
	fmt.Println(isIsomorphic(s, t))

	s, t = "foo", "bar"
	fmt.Println(isIsomorphic(s, t))

	s, t = "paper", "title"
	fmt.Println(isIsomorphic(s, t))
}
