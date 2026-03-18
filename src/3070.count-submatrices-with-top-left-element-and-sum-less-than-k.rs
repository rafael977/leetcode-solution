/*
 * @lc app=leetcode.cn id=3070 lang=rust
 * @lcpr version=30204
 *
 * [3070] Count Submatrices with Top-Left Element and Sum Less Than k
 */

// @lcpr-template-start
struct Solution;
// @lcpr-template-end
// @lc code=start
impl Solution {
    pub fn count_submatrices(grid: Vec<Vec<i32>>, k: i32) -> i32 {
        let mut count = 0;
        let m = grid.len();
        let n = grid[0].len();
        let mut cols = vec![0; n];
        for i in 0..m {
            let mut row_sum = 0;
            for j in 0..n {
                cols[j] += grid[i][j];
                row_sum += cols[j];
                if row_sum <= k {
                    count += 1;
                }
            }
        }

        count
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (vec![vec![7, 6, 3], vec![6, 6, 1]], 18, 4),
            (vec![vec![7, 2, 9], vec![1, 5, 0], vec![2, 6, 6]], 20, 6),
        ];
        for (i, (grid, k, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::count_submatrices(grid, k);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
