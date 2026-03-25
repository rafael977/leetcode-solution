/*
 * @lc app=leetcode.cn id=3546 lang=rust
 * @lcpr version=30204
 *
 * [3546] Equal Sum Grid Partition I
 */

// @lcpr-template-start
struct Solution;
// @lcpr-template-end
// @lc code=start
impl Solution {
    pub fn can_partition_grid(grid: Vec<Vec<i32>>) -> bool {
        let m = grid.len();
        let n = grid[0].len();
        let mut row_sum = vec![0i64; m];
        let mut col_sum = vec![0i64; n];
        let mut sum = 0i64;
        for i in 0..m {
            for j in 0..n {
                let x = grid[i][j] as i64;
                row_sum[i] += x;
                col_sum[j] += x;
                sum += x;
            }
        }
        if sum % 2 == 1 {
            return false;
        }
        let mut rs = 0;
        for s in row_sum {
            rs += s;
            if sum / 2 == rs {
                return true;
            }
        }
        let mut cs = 0;
        for s in col_sum {
            cs += s;
            if sum / 2 == cs {
                return true;
            }
        }

        false
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            // (vec![vec![1, 4], vec![2, 3]], true),
            // (vec![vec![1, 3], vec![2, 4]], false),
            // (vec![vec![14742, 71685, 59237, 27190]], true),
            (vec![vec![1, 1, 1]], false),
        ];
        for (i, (grid, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::can_partition_grid(grid);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {:#?}, got: {:#?}",
                i, expected, got
            )
        }
    }
}
