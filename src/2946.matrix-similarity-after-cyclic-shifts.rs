/*
 * @lc app=leetcode.cn id=2946 lang=rust
 * @lcpr version=30204
 *
 * [2946] Matrix Similarity After Cyclic Shifts
 */

// @lcpr-template-start
struct Solution;
// @lcpr-template-end
// @lc code=start
impl Solution {
    pub fn are_similar(mat: Vec<Vec<i32>>, k: i32) -> bool {
        let m = mat.len();
        let n = mat[0].len();
        let k = k % n as i32;
        fn check_row(row: &Vec<i32>, k: i32, factor: i32) -> bool {
            let n = row.len();
            for i in 0..n {
                let index_to_compare = ((i as i32 + k * factor + n as i32) % (n as i32)) as usize;
                if row[i] != row[index_to_compare] {
                    return false;
                }
            }
            true
        }

        for i in 0..m {
            if !check_row(&mat[i], k, if i % 2 == 0 { -1 } else { 1 }) {
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
            (vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]], 4, false),
            (
                vec![vec![1, 2, 1, 2], vec![5, 5, 5, 5], vec![6, 3, 6, 3]],
                2,
                true,
            ),
            (vec![vec![2, 2], vec![2, 2]], 3, true),
        ];
        for (i, (grid, k, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::are_similar(grid, k);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n4\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,1,2],[5,5,5,5],[6,3,6,3]]\n2\n
// @lcpr case=end

// @lcpr case=start
// [[2,2],[2,2]]\n3\n
// @lcpr case=end

 */
