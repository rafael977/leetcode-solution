/*
 * @lc app=leetcode.cn id=1886 lang=rust
 * @lcpr version=30204
 *
 * [1886] Determine Whether Matrix Can Be Obtained By Rotation
 */

// @lcpr-template-start
struct Solution;
// @lcpr-template-end
// @lc code=start
impl Solution {
    pub fn find_rotation(mat: Vec<Vec<i32>>, target: Vec<Vec<i32>>) -> bool {
        let mut ok = [true; 4];
        let n = mat.len();
        for i in 0..n {
            for j in 0..n {
                if mat[i][j] != target[i][j] {
                    ok[0] = false;
                }
                if mat[i][j] != target[j][n - 1 - i] {
                    ok[1] = false;
                }
                if mat[i][j] != target[n - 1 - i][n - 1 - j] {
                    ok[2] = false;
                }
                if mat[i][j] != target[n - 1 - j][i] {
                    ok[3] = false;
                }
                if ok.iter().all(|b| *b == false) {
                    return false;
                }
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
            (
                vec![vec![0, 1], vec![1, 0]],
                vec![vec![1, 0], vec![0, 1]],
                true,
            ),
            (
                vec![vec![0, 1], vec![1, 1]],
                vec![vec![1, 0], vec![0, 1]],
                false,
            ),
            (
                vec![vec![0, 0, 0], vec![0, 1, 0], vec![1, 1, 1]],
                vec![vec![1, 1, 1], vec![0, 1, 0], vec![0, 0, 0]],
                true,
            ),
        ];
        for (i, (mat, target, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::find_rotation(mat, target);
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
// [[0,1],[1,0]]\n[[1,0],[0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[1,1]]\n[[1,0],[0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0],[0,1,0],[1,1,1]]\n[[1,1,1],[0,1,0],[0,0,0]]\n
// @lcpr case=end

 */
