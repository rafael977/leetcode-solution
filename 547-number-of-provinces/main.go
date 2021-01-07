package main

import "fmt"

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	count := 0
	for i := range isConnected {
		if visited[i] {
			continue
		}
		count++
		dfs(isConnected, i, visited)
	}

	return count
}
func dfs(isConnected [][]int, i int, visited []bool) {
	visited[i] = true
	for j := range isConnected[i] {
		if isConnected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(isConnected, j, visited)
		}
	}
}

func main() {
	g := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(g))

	g = [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(g))

	g = [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}
	fmt.Println(findCircleNum(g))
}
