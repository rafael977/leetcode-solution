package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 2 {
		return false
	}

	p1, p2 := coordinates[0], coordinates[1]
	a, b, base := p1[1]-p2[1], p1[0]*p2[1]-p1[1]*p2[0], p1[0]-p2[0]
	for _, v := range coordinates[2:] {
		if a*v[0]+b != v[1]*base {
			return false
		}
	}

	return true
}

func main() {
	pts := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	fmt.Println(checkStraightLine(pts))
}
