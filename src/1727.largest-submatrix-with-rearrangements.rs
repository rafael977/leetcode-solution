/*
 * @lc app=leetcode.cn id=1727 lang=rust
 * @lcpr version=30204
 *
 * [1727] Largest Submatrix With Rearrangements
 */

// @lcpr-template-start

// @lcpr-template-end
struct Solution;

// @lc code=start
use std::cmp;

impl Solution {
    pub fn largest_submatrix(matrix: Vec<Vec<i32>>) -> i32 {
        let n = matrix[0].len();
        let mut height = vec![0; n];
        let mut ans = 0;

        for row in matrix.iter() {
            for (j, v) in row.iter().enumerate() {
                if *v == 0 {
                    height[j] = 0;
                } else {
                    height[j] += 1;
                }
            }

            let mut hs = height.clone();
            hs.sort();
            for (j, h) in hs.into_iter().enumerate() {
                ans = cmp::max(ans, (n - j) * h)
            }
        }

        ans as i32
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (vec![vec![0, 0, 1], vec![1, 1, 1], vec![1, 0, 1]], 4),
            (vec![vec![1, 0, 1, 0, 1]], 3),
            (vec![vec![1, 1, 0], vec![1, 0, 1]], 2),
        ];
        for (i, (matrix, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::largest_submatrix(matrix);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
