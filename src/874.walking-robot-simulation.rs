/*
 * @lc app=leetcode.cn id=874 lang=rust
 * @lcpr version=30204
 *
 * [874] Walking Robot Simulation
 */

// @lcpr-template-start
struct Solution;
// @lcpr-template-end
// @lc code=start
use std::{cmp, collections::HashSet};

impl Solution {
    pub fn robot_sim(commands: Vec<i32>, obstacles: Vec<Vec<i32>>) -> i32 {
        let dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)];
        let obstacles = obstacles
            .into_iter()
            .map(|p| (p[0], p[1]))
            .collect::<HashSet<_>>();
        let mut k = 0usize;
        let mut pos = (0, 0);
        let mut ans = 0;
        for c in commands {
            if c == -1 {
                k = (k + 1) % 4;
                continue;
            }
            if c == -2 {
                k = (k + 3) % 4;
                continue;
            }
            let dir = &dirs[k];
            for _ in 0..c {
                if obstacles.contains(&(pos.0 + dir.0, pos.1 + dir.1)) {
                    break;
                }
                pos.0 += dir.0;
                pos.1 += dir.1;
            }
            ans = cmp::max(ans, pos.0 * pos.0 + pos.1 * pos.1)
        }

        ans
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (vec![4, -1, 3], vec![], 25),
            (vec![4, -1, 4, -2, 4], vec![vec![2, 4]], 65),
            (vec![6, -1, -1, 6], vec![], 36),
        ];
        for (i, (commands, obstacles, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::robot_sim(commands, obstacles);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {:#?}, got: {:#?}",
                i, expected, got
            )
        }
    }
}
