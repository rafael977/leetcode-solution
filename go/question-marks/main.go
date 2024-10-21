package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func QuestionsMarks(str string) string {
	ln := len(str)
	sArr := strings.Split(str, "")
	r := regexp.MustCompile(`[\w]*\?[\w]*\?[\w]*\?[\w]*`)
	ans := "false"
	for i := 0; i < ln; i++ {
		vi, err := strconv.Atoi(sArr[i])
		if err != nil {
			continue
		}
		for j := i; j < ln-i; j++ {
			vj, err := strconv.Atoi(sArr[j])
			if err != nil || vi+vj != 10 {
				continue
			}

			ans = "true"
			if !r.MatchString(strings.Join(sArr[i+1:j], "")) {
				return "false"
			}
		}
	}
	// code goes here
	return ans

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(QuestionsMarks("aa6?9"))

	fmt.Println(QuestionsMarks("acc?7??sss?3rr1??????5"))
}
