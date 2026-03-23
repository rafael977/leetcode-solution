/*
 * @lc app=leetcode.cn id=1594 lang=rust
 * @lcpr version=30204
 *
 * [1594] Maximum Non Negative Product in a Matrix
 */

// @lcpr-template-start
struct Solution;
// @lcpr-template-end
// @lc code=start
use std::cmp;

const MOD: i64 = 1_000_000_007;

impl Solution {
    pub fn max_product_path(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();

        let mut mem = vec![vec![vec![0; 2]; n]; m];
        mem[0][0][0] = grid[0][0] as i64;
        mem[0][0][1] = grid[0][0] as i64;
        for i in 1..m {
            let x = grid[i][0] as i64;
            mem[i][0][0] = cmp::min(mem[i - 1][0][0] * x, mem[i - 1][0][1] * x);
            mem[i][0][1] = cmp::max(mem[i - 1][0][0] * x, mem[i - 1][0][1] * x);
        }
        for j in 1..n {
            let x = grid[0][j] as i64;
            mem[0][j][0] = cmp::min(mem[0][j - 1][0] * x, mem[0][j - 1][1] * x);
            mem[0][j][1] = cmp::max(mem[0][j - 1][0] * x, mem[0][j - 1][1] * x);
        }
        for i in 1..m {
            for j in 1..n {
                let x = grid[i][j] as i64;
                mem[i][j][0] = cmp::min(
                    cmp::min(mem[i - 1][j][0] * x, mem[i - 1][j][1] * x),
                    cmp::min(mem[i][j - 1][0] * x, mem[i][j - 1][1] * x),
                );
                mem[i][j][1] = cmp::max(
                    cmp::max(mem[i - 1][j][0] * x, mem[i - 1][j][1] * x),
                    cmp::max(mem[i][j - 1][0] * x, mem[i][j - 1][1] * x),
                );
            }
        }

        let ans = mem[m - 1][n - 1][1];
        if ans < 0 { -1 } else { (ans % MOD) as i32 }
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
                vec![vec![-1, -2, -3], vec![-2, -3, -3], vec![-3, -3, -2]],
                -1,
            ),
            (vec![vec![1, -2, 1], vec![1, -2, 1], vec![3, -4, 1]], 8),
            (vec![vec![1, 3], vec![0, -4]], 0),
            (
                vec![
                    vec![2, 1, 3, 0, -3, 3, -4, 4, 0, -4],
                    vec![-4, -3, 2, 2, 3, -3, 1, -1, 1, -2],
                    vec![-2, 0, -4, 2, 4, -3, -4, -1, 3, 4],
                    vec![-1, 0, 1, 0, -3, 3, -2, -3, 1, 0],
                    vec![0, -1, -2, 0, -3, -4, 0, 3, -2, -2],
                    vec![-4, -2, 0, -1, 0, -3, 0, 4, 0, -3],
                    vec![-3, -4, 2, 1, 0, -4, 2, -4, -1, -3],
                    vec![3, -2, 0, -4, 1, 0, 1, -3, -1, -1],
                    vec![3, -4, 0, 2, 0, -2, 2, -4, -2, 4],
                    vec![0, 4, 0, -3, -4, 3, 3, -1, -2, -2],
                ],
                19215865,
            ),
            (
                vec![
                    vec![1, -1, 2, 1, -1, 0, 0, 4, 3, 2, 0, -2, -2],
                    vec![-2, 3, 3, -1, -1, 0, 0, -2, 4, -3, 3, 0, 0],
                    vec![-4, -1, -1, -2, 2, -1, -2, -2, 0, 3, -1, -4, 1],
                    vec![-3, 4, -3, 0, -3, 1, -3, 1, 4, 4, -4, -4, -2],
                    vec![3, -3, 1, 0, -1, -4, -4, -4, 3, 2, 2, 3, 3],
                    vec![2, -1, -1, -4, -3, -3, 4, 2, 3, 4, 4, -4, 0],
                    vec![4, -1, 2, -3, -1, -1, -3, -4, 4, 4, 4, -3, -1],
                    vec![-3, -4, 4, -2, -1, 2, 3, -1, 2, 3, 4, 4, -4],
                    vec![-3, -1, -2, 1, 1, -1, -3, -4, -3, 1, -3, 3, -4],
                    vec![2, 4, 4, 4, -3, -3, 1, -1, 3, 4, -1, 1, 4],
                    vec![2, -2, 0, 4, -1, 0, -2, 4, -4, 0, 0, 2, -3],
                    vec![1, 1, -3, 0, -4, -4, -4, -4, 0, -1, -4, -1, 0],
                    vec![3, -1, -3, -3, -3, -2, -1, 4, -1, -2, 4, 2, 3],
                ],
                459630706,
            ),
        ];
        for (i, (grid, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::max_product_path(grid);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
