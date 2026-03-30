/*
 * @lc app=leetcode.cn id=2840 lang=rust
 * @lcpr version=30204
 *
 * [2840] Check if Strings Can be Made Equal With Operations II
 */

// @lcpr-template-start
struct Solution;
// @lcpr-template-end
// @lc code=start
impl Solution {
    pub fn check_strings(s1: String, s2: String) -> bool {
        let s1 = s1.into_bytes();
        let s2 = s2.into_bytes();
        let mut c_count = [[0; 26]; 2];
        for i in 0..s1.len() {
            c_count[i % 2][(s1[i] - b'a') as usize] += 1;
            c_count[i % 2][(s2[i] - b'a') as usize] -= 1;
        }

        c_count[0].iter().all(|&c| c == 0) && c_count[1].iter().all(|&c| c == 0)
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![("abcdba", "cabdab", true), ("abe", "bea", false)];
        for (i, (s1, s2, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::check_strings(String::from(s1), String::from(s2));
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {:#?}, got: {:#?}",
                i, expected, got
            )
        }
    }
}

/*
// @lcpr case=start
// "abcdba"\n"cabdab"\n
// @lcpr case=end

// @lcpr case=start
// "abe"\n"bea"\n
// @lcpr case=end

 */
