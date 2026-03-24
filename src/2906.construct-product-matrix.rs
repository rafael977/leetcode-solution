/*
 * @lc app=leetcode.cn id=2906 lang=rust
 * @lcpr version=30204
 *
 * [2906] Construct Product Matrix
 */

// @lcpr-template-start
struct Solution;
// @lcpr-template-end
// @lc code=start
const MOD: i64 = 12345;

impl Solution {
    pub fn construct_product_matrix(grid: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let n = grid.len();
        let m = grid[0].len();
        let mut p = vec![vec![0; m]; n];

        let mut suf = 1;
        for i in (0..n).rev() {
            for j in (0..m).rev() {
                p[i][j] = suf as i32;
                suf = suf * grid[i][j] as i64 % MOD;
            }
        }

        let mut pre = 1;
        for i in 0..n {
            for j in 0..m {
                p[i][j] = (p[i][j] as i64 * pre % MOD) as i32;
                pre = pre * grid[i][j] as i64 % MOD;
            }
        }

        p
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (vec![vec![1, 2], vec![3, 4]], vec![vec![24, 12], vec![8, 6]]),
            (
                vec![vec![12345], vec![2], vec![1]],
                vec![vec![2], vec![0], vec![0]],
            ),
        ];
        for (i, (grid, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::construct_product_matrix(grid);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {:#?}, got: {:#?}",
                i, expected, got
            )
        }
    }
}
