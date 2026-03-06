/*
 * @lc app=leetcode.cn id=1784 lang=rust
 * @lcpr version=30204
 *
 * [1784] Check if Binary String Has at Most One Segment of Ones
 */

// @lcpr-template-start

// @lcpr-template-end
struct Solution;

// @lc code=start
impl Solution {
    pub fn check_ones_segment(s: String) -> bool {
        let mut found_one = false;
        for c in s.into_bytes().iter().rev() {
            if *c == b'1' {
                found_one = true;
            }
            if *c == b'0' && found_one {
                return false;
            }
        }

        true
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (String::from("1001"), false),
            (String::from("110"), true),
            (String::from("10"), true),
        ];
        for (i, (s, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::check_ones_segment(s);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
