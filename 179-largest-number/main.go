package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type IntString []int

func (x IntString) Len() int {
	return len(x)
}

func (x IntString) Less(i, j int) bool {
	return fmt.Sprintf("%d%d", x[i], x[j]) > fmt.Sprintf("%d%d", x[j], x[i])
}

func (x IntString) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func largestNumber(nums []int) string {
	data := IntString(nums)
	sort.Sort(data)

	if data[0] == 0 {
		return "0"
	}

	numsText := make([]string, len(data))
	for i := range data {
		numsText[i] = strconv.Itoa(data[i])
	}

	return strings.Join(numsText, "")
}

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
