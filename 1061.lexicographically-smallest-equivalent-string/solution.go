package lexicographicallysmallestequivalentstring

import "strings"

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := make([]byte, 26)
	for i := range parent {
		parent[i] = byte(i)
	}

	for i := range s1 {
		union(parent, s1[i]-'a', s2[i]-'a')
	}

	sb := strings.Builder{}
	for i := range baseStr {
		c := find(parent, baseStr[i]-'a') + 'a'
		sb.WriteByte(c)
	}

	return sb.String()
}

func find(parent []byte, i byte) byte {
	if i == parent[i] {
		return i
	}
	p := find(parent, parent[i])
	parent[i] = p
	return p
}

func union(parent []byte, i, j byte) {
	pi, pj := find(parent, i), find(parent, j)
	if pi < pj {
		parent[pj] = pi
	} else {
		parent[pi] = pj
	}
}
