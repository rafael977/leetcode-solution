package leetcodesolution

import "strings"

/*
 * @lc app=leetcode id=273 lang=golang
 *
 * [273] Integer to English Words
 */

// @lc code=start
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	numWordMap := map[int]string{
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
		20: "Twenty",
		30: "Thirty",
		40: "Forty",
		50: "Fifty",
		60: "Sixty",
		70: "Seventy",
		80: "Eighty",
		90: "Ninety",
	}
	order := map[int]string{
		1:  "Thousand",
		2:  "Million",
		3:  "Billion",
		4:  "Trillion",
		5:  "Quadrillion",
		6:  "Quintillion",
		7:  "Sextillion",
		8:  "Septillion",
		9:  "Octillion",
		10: "Nonillion",
	}

	convertWithinHundred := func(i int) string {
		if i == 0 {
			return ""
		}
		if i <= 20 {
			return numWordMap[i]
		}
		tens, ones := i/10*10, i%10
		if ones == 0 {
			return numWordMap[tens]
		}
		return numWordMap[tens] + " " + numWordMap[ones]
	}

	convert3Digits := func(i int) string {
		if i < 100 {
			return convertWithinHundred(i)
		}
		hundreds := i / 100
		digit2 := convertWithinHundred(i % 100)
		if digit2 == "" {
			return numWordMap[hundreds] + " Hundred"
		}
		return numWordMap[hundreds] + " Hundred " + digit2
	}

	ans := ""
	o := 0
	for num > 0 {
		x := num % 1000
		num = num / 1000
		if x == 0 {
			o++
			continue
		}
		if o == 0 {
			ans = convert3Digits(x)
		} else {
			ans = convert3Digits(x) + " " + order[o] + " " + ans
		}
		o++
	}

	return strings.Trim(ans, " ")
}

// @lc code=end
