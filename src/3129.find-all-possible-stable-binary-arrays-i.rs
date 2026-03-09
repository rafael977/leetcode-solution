/*
 * @lc app=leetcode.cn id=3129 lang=rust
 * @lcpr version=30204
 *
 * [3129] Find All Possible Stable Binary Arrays I
 */

// @lcpr-template-start

// @lcpr-template-end
struct Solution;
// @lc code=start
use std::collections::HashMap;

const MOD: i32 = 1_000_000_007;

impl Solution {
    pub fn number_of_stable_arrays(zero: i32, one: i32, limit: i32) -> i32 {
        fn dfs(
            i: u32,
            j: u32,
            k: u32,
            z: &i32,
            o: &i32,
            l: &i32,
            memo: &mut HashMap<(u32, u32, u32), i32>,
        ) -> i32 {
            if i == 0 {
                if k == 1 && j <= *l as u32 {
                    return 1;
                }
                return 0;
            }
            if j == 0 {
                if k == 0 && i <= *l as u32 {
                    return 1;
                }
                return 0;
            }

            if memo.contains_key(&(i, j, k)) {
                return memo[&(i, j, k)];
            }
            if k == 0 {
                let mut v =
                    (dfs(i - 1, j, 0, z, o, l, memo) + dfs(i - 1, j, 1, z, o, l, memo)) % MOD;
                if i > *l as u32 {
                    v = (v + MOD - dfs(i - *l as u32 - 1, j, 1, z, o, l, memo)) % MOD
                }

                memo.insert((i, j, k), v);
                return v;
            }

            let mut v = (dfs(i, j - 1, 0, z, o, l, memo) + dfs(i, j - 1, 1, z, o, l, memo)) % MOD;
            if j > *l as u32 {
                v = (v + MOD - dfs(i, j - *l as u32 - 1, 0, z, o, l, memo)) % MOD;
            }
            memo.insert((i, j, k), v);
            return v;
        }

        let mut memo: HashMap<(u32, u32, u32), i32> = HashMap::new();
        (dfs(zero as u32, one as u32, 0, &zero, &one, &limit, &mut memo)
            + dfs(zero as u32, one as u32, 1, &zero, &one, &limit, &mut memo))
            % MOD
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
