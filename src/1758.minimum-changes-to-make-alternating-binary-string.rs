/*
 * @lc app=leetcode.cn id=1758 lang=rust
 * @lcpr version=30204
 *
 * [1758] 生成交替二进制字符串的最少操作数
 */

// @lcpr-template-start

// @lcpr-template-end
struct Solution;
// @lc code=start
use std::cmp;

impl Solution {
    pub fn min_operations(s: String) -> i32 {
        let diff = |s: &String, odd_char: u8, even_char: u8| -> i32 {
            let mut cnt = 0;
            for (i, c) in s.as_bytes().iter().enumerate() {
                if i % 2 == 0 && *c != even_char {
                    cnt += 1;
                }
                if i % 2 == 1 && *c != odd_char {
                    cnt += 1;
                }
            }
            cnt
        };

        let v1 = diff(&s, b'0', b'1');
        let v2 = diff(&s, b'1', b'0');

        cmp::min(v1, v2)
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (String::from("0100"), 1),
            (String::from("10"), 0),
            (String::from("1111"), 2),
        ];
        for (i, (s, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::min_operations(s);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
