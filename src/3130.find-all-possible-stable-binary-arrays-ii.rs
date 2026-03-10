/*
 * @lc app=leetcode.cn id=3130 lang=rust
 * @lcpr version=30204
 *
 * [3130] Find All Possible Stable Binary Arrays II
 */

// @lcpr-template-start

// @lcpr-template-end
struct Solution;

// @lc code=start
use std::cmp;

const MOD: i32 = 1_000_000_007;

impl Solution {
    pub fn number_of_stable_arrays(zero: i32, one: i32, limit: i32) -> i32 {
        let zero = zero as usize;
        let one = one as usize;
        let limit = limit as usize;
        let mut memo = vec![vec![vec![0; 2]; one + 1]; zero + 1];
        for i in 0..=cmp::min(limit, zero) {
            memo[i][0][0] = 1;
        }
        for j in 0..=cmp::min(limit, one) {
            memo[0][j][1] = 1;
        }
        for i in 1..=zero {
            for j in 1..=one {
                memo[i][j][0] = (memo[i - 1][j][0] + memo[i - 1][j][1]) % MOD;
                if i > limit {
                    memo[i][j][0] = (memo[i][j][0] + MOD - memo[i - limit - 1][j][1]) % MOD; // +MOD to prevent negative
                }
                memo[i][j][1] = (memo[i][j - 1][0] + memo[i][j - 1][1]) % MOD;
                if j > limit {
                    memo[i][j][1] = (memo[i][j][1] + MOD - memo[i][j - limit - 1][0]) % MOD;
                }
            }
        }

        (memo[zero][one][0] + memo[zero][one][1]) % MOD
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![(1, 1, 2, 2), (1, 2, 1, 1), (3, 3, 2, 14)];
        for (i, (zero, one, limit, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::number_of_stable_arrays(zero, one, limit);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
