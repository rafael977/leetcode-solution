package asteroidcollision

/*
 * @lc app=leetcode id=735 lang=golang
 *
 * [735] Asteroid Collision
 */

// @lc code=start
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, v := range asteroids {
		if len(stack) == 0 || v > 0 {
			stack = append(stack, v)
			continue
		}

		explode := false
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top < 0 {
				break
			}
			if top <= -v {
				stack = stack[:len(stack)-1]
				if top == -v {
					explode = true
					break
				}
			} else {
				explode = true
				break
			}
		}
		if !explode {
			stack = append(stack, v)
		}
	}
	return stack
}

// @lc code=end
