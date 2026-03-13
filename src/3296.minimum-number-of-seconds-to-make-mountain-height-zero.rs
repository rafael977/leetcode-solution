/*
 * @lc app=leetcode.cn id=3296 lang=rust
 * @lcpr version=30204
 *
 * [3296] Minimum Number of Seconds to Make Mountain Height Zero
 */

// @lcpr-template-start

// @lcpr-template-end
struct Solution;

// @lc code=start
use std::cmp::Reverse;
use std::collections::BinaryHeap;

impl Solution {
    pub fn min_number_of_seconds(mountain_height: i32, worker_times: Vec<i32>) -> i64 {
        let mut pq = BinaryHeap::new();
        for (i, w) in worker_times.iter().enumerate() {
            pq.push(Reverse((*w as i64, i, (*w + *w) as i64)));
        }
        for _ in 0..mountain_height - 1 {
            let Reverse((w, i, nw)) = pq.pop().unwrap();
            pq.push(Reverse((w + nw, i, nw + worker_times[i] as i64)));
        }

        let Reverse(last) = pq.pop().unwrap();
        last.0 as i64
    }
}
// @lc code=end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (4, vec![2, 1, 1], 3),
            (10, vec![3, 2, 2, 4], 12),
            (5, vec![1], 15),
            (5, vec![1, 5], 10),
            (100000, vec![1], 5000050000),
        ];
        for (i, (mountain_height, worker_times, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::min_number_of_seconds(mountain_height, worker_times);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
