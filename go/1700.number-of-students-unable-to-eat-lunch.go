package leetcodesolution

/*
 * @lc app=leetcode id=1700 lang=golang
 *
 * [1700] Number of Students Unable to Eat Lunch
 */

// @lc code=start
func countStudents(students []int, sandwiches []int) int {
	i := 0
	shift := 0
	for len(students) > 0 {
		if students[0] == sandwiches[i] {
			students = students[1:]
			i++
			shift = 0
			continue
		}
		students = append(students, students[0])
		students = students[1:]
		shift++
		if shift == len(students) {
			break
		}
	}
	return len(students)
}

// @lc code=end
