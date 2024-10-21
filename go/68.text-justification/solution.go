package textjustification

import "strings"

/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	ans := make([]string, 0)
	i := 0
	for i < len(words) {
		row := make([]string, 0)
		width := 0
		for i < len(words) &&
			width+len(row)+len(words[i]) <= maxWidth {
			row = append(row, words[i])
			width += len(words[i])
			i++
		}
		if i == len(words) {
			ans = append(ans, buildLast(row, maxWidth))
		} else {
			ans = append(ans, build(row, width, maxWidth))
		}
	}
	return ans
}

func build(row []string, width, maxWidth int) string {
	sb := strings.Builder{}
	if len(row) == 1 {
		sb.WriteString(row[0])
		for sb.Len() < maxWidth {
			sb.WriteRune(' ')
		}
		return sb.String()
	}

	sp := (maxWidth - width) / (len(row) - 1)
	r := (maxWidth - width) % (len(row) - 1)
	for i, v := range row {
		if i > 0 {
			n := sp
			if i <= r {
				n++
			}
			for i := 0; i < n; i++ {
				sb.WriteRune(' ')
			}
		}
		sb.WriteString(v)
	}
	return sb.String()
}

func buildLast(row []string, maxWidth int) string {
	sb := strings.Builder{}
	for _, v := range row {
		sb.WriteString(v)
		if sb.Len() < maxWidth {
			sb.WriteRune(' ')
		}
	}
	for sb.Len() < maxWidth {
		sb.WriteRune(' ')
	}
	return sb.String()
}

// @lc code=end
