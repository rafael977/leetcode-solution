/*
 * @lc app=leetcode.cn id=3643 lang=rust
 * @lcpr version=30204
 *
 * [3643] Flip Square Submatrix Vertically
 */

// @lcpr-template-start
struct Solution;
// @lcpr-template-end
// @lc code=start
impl Solution {
    pub fn reverse_submatrix(mut grid: Vec<Vec<i32>>, x: i32, y: i32, k: i32) -> Vec<Vec<i32>> {
        let x = x as usize;
        let y = y as usize;
        let k = k as usize;

        for j in y..y + k {
            let (mut s, mut t) = (x, x + k - 1);
            while s < t {
                (grid[s][j], grid[t][j]) = (grid[t][j], grid[s][j]);
                s += 1;
                t -= 1;
            }
        }

        grid
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (
                vec![
                    vec![1, 2, 3, 4],
                    vec![5, 6, 7, 8],
                    vec![9, 10, 11, 12],
                    vec![13, 14, 15, 16],
                ],
                1,
                0,
                3,
                vec![
                    vec![1, 2, 3, 4],
                    vec![13, 14, 15, 8],
                    vec![9, 10, 11, 12],
                    vec![5, 6, 7, 16],
                ],
            ),
            (
                vec![vec![3, 4, 2, 3], vec![2, 3, 4, 2]],
                0,
                2,
                2,
                vec![vec![3, 4, 4, 2], vec![2, 3, 2, 3]],
            ),
        ];
        for (i, (grid, x, y, k, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::reverse_submatrix(grid, x, y, k);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {:#?}, got: {:#?}",
                i, expected, got
            )
        }
    }
}
