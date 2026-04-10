/*
 * @lc app=leetcode.cn id=3740 lang=rust
 * @lcpr version=30204
 *
 * [3740] Minimum Distance Between Three Equal Elements I
 */
struct Solution;

// @lc code=start
use std::{cmp, collections::HashMap};

impl Solution {
    pub fn minimum_distance(nums: Vec<i32>) -> i32 {
        let mut map: HashMap<i32, Vec<usize>> = HashMap::new();
        for (i, n) in nums.into_iter().enumerate() {
            if map.contains_key(&n) {
                map.get_mut(&n).unwrap().push(i);
                continue;
            }
            map.insert(n, vec![i]);
        }

        let mut ans = i32::MAX;
        for (_, v) in map {
            if v.len() < 3 {
                continue;
            }
            for i in 2..v.len() {
                ans = cmp::min(ans, (v[i] - v[i - 2]) as i32 * 2);
            }
        }
        if ans == i32::MAX { -1 } else { ans }
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (vec![1, 2, 1, 1, 3], 6),
            (vec![1, 1, 2, 3, 2, 1, 2], 8),
            (vec![1], -1),
        ];
        for (i, (nums, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::minimum_distance(nums);
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
// [1,2,1,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2,3,2,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

 */
