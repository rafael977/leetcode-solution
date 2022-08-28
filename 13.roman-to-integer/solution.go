package romantointeger

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
var nmap map[string]int = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) (ans int) {
	for i := 0; i < len(s); {
		if i != len(s)-1 {
			n := s[i : i+2]
			if v, ok := nmap[n]; ok {
				ans += v
				i += 2
				continue
			}
		}
		n := s[i : i+1]
		ans += nmap[n]
		i++
	}
	return
}

// @lc code=end
