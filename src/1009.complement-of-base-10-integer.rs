/*
 * @lc app=leetcode.cn id=1009 lang=rust
 * @lcpr version=30204
 *
 * [1009] Complement of Base 10 Integer
 */

// @lcpr-template-start

// @lcpr-template-end
struct Solution;

// @lc code=start
impl Solution {
    pub fn bitwise_complement(n: i32) -> i32 {
        let mut mask = 2;
        while mask <= n {
            mask = mask << 1;
        }

        n ^ (mask - 1)
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![(5, 2), (7, 0), (10, 5), (0, 1), (1, 0), (2, 1)];
        for (i, (n, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::bitwise_complement(n);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {:#?}, got: {:#?}",
                i, expected, got
            )
        }
    }
}
