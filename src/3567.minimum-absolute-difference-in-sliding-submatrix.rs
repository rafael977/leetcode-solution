/*
 * @lc app=leetcode.cn id=3567 lang=rust
 * @lcpr version=30204
 *
 * [3567] Minimum Absolute Difference in Sliding Submatrix
 */

// @lcpr-template-start
struct Solution;

// @lcpr-template-end
// @lc code=start
use std::{cmp, collections::HashSet, i32::MAX};

impl Solution {
    pub fn min_abs_diff(grid: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        let m = grid.len();
        let n = grid[0].len();
        let k = k as usize;
        let mut ans = vec![vec![0; n - k + 1]; m - k + 1];
        if k == 1 {
            return ans;
        }
        for i in 0..(m - k + 1) {
            for j in 0..(n - k + 1) {
                let mut sub = HashSet::new();
                for ii in i..(i + k) {
                    for jj in j..(j + k) {
                        sub.insert(grid[ii][jj]);
                    }
                }
                if sub.len() == 1 {
                    break;
                }
                let mut sub: Vec<i32> = sub.into_iter().collect();
                sub.sort();
                let mut min = MAX;
                for x in 0..sub.len() - 1 {
                    min = cmp::min(min, sub[x + 1] - sub[x]);
                }
                ans[i][j] = min;
            }
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
            (vec![vec![1, 8], vec![3, -2]], 2, vec![vec![2]]),
            (vec![vec![3, -1]], 1, vec![vec![0, 0]]),
            (vec![vec![1, -2, 3], vec![2, 3, 5]], 2, vec![vec![1, 2]]),
        ];
        for (i, (grid, k, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::min_abs_diff(grid, k);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {:#?}, got: {:#?}",
                i, expected, got
            )
        }
    }
}
