/*
 * @lc app=leetcode.cn id=1980 lang=rust
 * @lcpr version=30204
 *
 * [1980] Find Unique Binary String
 */

// @lcpr-template-start

// @lcpr-template-end
struct Solution;

// @lc code=start
use std::collections::HashSet;
impl Solution {
    pub fn find_different_binary_string(nums: Vec<String>) -> String {
        let n = nums.len();
        let mut values = HashSet::new();
        for num in &nums {
            let v = u32::from_str_radix(num, 2).unwrap();
            values.insert(v);
        }

        let mut i = 0u32;
        while i < (1 << n) {
            if !values.contains(&i) {
                break;
            }
            i += 1;
        }

        let mut r = format!("{:b}", i);
        while r.len() < n {
            r = format!("0{}", r);
        }

        r
    }
}
// @lc code=end

// impl Solution {
//     pub fn find_different_binary_string(nums: Vec<String>) -> String {
//         let n = (&nums[0]).len();
//         fn base10_value<'a>(num: &'a str, cache: &mut HashMap<&'a str, u32>) -> u32 {
//             if num == "0" {
//                 return 0;
//             }
//             if num == "1" {
//                 return 1;
//             }
//             if cache.contains_key(num) {
//                 return cache[num];
//             }
//             let v = (base10_value(&num[..num.len() - 1], cache) << 1)
//                 + match num.as_bytes()[num.len() - 1] {
//                     x if x == b'1' => 1,
//                     _ => 0,
//                 };
//             cache.insert(num, v);
//             v
//         }

//         let mut values = HashMap::new();
//         for num in nums {
//             let mut cache = HashMap::new();
//             values.insert(base10_value(&num.as_str(), &mut cache), true);
//         }

//         for i in 0u32..(1 << n) {
//             if !values.contains_key(&i) {
//                 let mut v = i;
//                 let mut r = vec![];
//                 loop {
//                     if v == 0 {
//                         for _ in r.len()..n {
//                             r.push("0".to_string());
//                         }
//                         return String::from_iter(r.into_iter().rev());
//                     }
//                     r.push(format!("{}", v & 1));
//                     v = v >> 1;
//                 }
//             }
//         }

//         String::new()
//     }
// }

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            // (vec!["01".to_string(), "10".to_string()], "00".to_string()),
            // (vec!["00".to_string(), "01".to_string()], "10".to_string()),
            // (
            //     vec!["111".to_string(), "011".to_string(), "001".to_string()],
            //     "000".to_string(),
            // ),
            (vec!["00".to_string(), "10".to_string()], "01".to_string()),
        ];
        for (i, (nums, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::find_different_binary_string(nums);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
