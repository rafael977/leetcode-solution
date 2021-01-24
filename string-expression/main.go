package main

import "fmt"

var expToNum map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}
var numToExp map[int]string = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
	0: "zero",
}

func StringExpression(str string) string {
	start := 0
	result := 0
	num := 0
	op := ""
	for i := 0; i <= len(str); i++ {
		fmt.Println("i", i)
		exp := str[start:i]
		if n, ok := expToNum[exp]; ok {
			num = num*10 + n
			start = i
		} else if exp == "plus" || exp == "minus" {
			if op == "" {
				result = num
			} else if op == "plus" {
				result = result + num
			} else if op == "minus" {
				result = result - num
			}
			num = 0
			op = exp
			start = i
		}
	}
	if op == "plus" {
		result += num
	} else if op == "minus" {
		result -= num
	}

	fmt.Println(result)
	return toExpression(result)
}

func toExpression(i int) string {
	if i == 0 {
		return "zero"
	}
	s := ""
	nStr := ""
	if i < 0 {
		s += "negative"
		i = -i
	}
	for i > 0 {
		nStr = fmt.Sprintf("%s%s", numToExp[i%10], nStr)
		i = i / 10
	}

	return fmt.Sprintf("%s%s", s, nStr)
}

func main() {
	fmt.Println(StringExpression("onezeropluseight"))
	fmt.Println(StringExpression("oneminusoneone"))
}
