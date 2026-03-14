/*
 * @lc app=leetcode.cn id=1415 lang=rust
 * @lcpr version=30204
 *
 * [1415] The k-th Lexicographical String of All Happy Strings of Length n
 */

// @lcpr-template-start

// @lcpr-template-end
struct Solution;

// @lc code=start
impl Solution {
    pub fn get_happy_string(n: i32, k: i32) -> String {
        let per_tree = 2i32.pow((n - 1) as u32);
        if k > 3 * per_tree {
            return String::from("");
        }

        let mut parent_char = match (k - 1) / per_tree {
            q if q == 0 => 'a',
            q if q == 1 => 'b',
            q if q == 2 => 'c',
            _ => panic!(),
        };

        let mut result = String::new();
        result.push(parent_char);

        let mut position = ((k - 1) % per_tree) + 1;
        let mut max_lots = per_tree;
        for _ in 1..n {
            let ch = match parent_char {
                c if position <= max_lots / 2 && c == 'a' => 'b',
                c if position <= max_lots / 2 && (c == 'b' || c == 'c') => 'a',
                c if position > max_lots / 2 && (c == 'a' || c == 'b') => 'c',
                c if position > max_lots / 2 && c == 'c' => 'b',
                _ => panic!(),
            };
            result.push(ch);
            parent_char = ch;
            max_lots = max_lots / 2;
            if position > max_lots {
                position -= max_lots;
            }
        }

        result
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (1, 3, "c"),
            (1, 4, ""),
            (3, 9, "cab"),
            (10, 100, "abacbabacb"),
            (7, 64, "acbcbcb"),
        ];
        for (i, (n, k, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::get_happy_string(n, k);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
