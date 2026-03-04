struct Solution;
// @lc start
impl Solution {
    pub fn num_special(mat: Vec<Vec<i32>>) -> i32 {
        let m = mat.len();
        let n = mat[0].len();
        let mut row_count = vec![0; m];
        let mut col_count = vec![0; n];
        for i in 0..m {
            for j in 0..n {
                if mat[i][j] == 1 {
                    row_count[i] += 1;
                    col_count[j] += 1;
                }
            }
        }

        let mut r = 0;
        for i in 0..m {
            for j in 0..n {
                if mat[i][j] == 1 && row_count[i] == 1 && col_count[j] == 1 {
                    r += 1;
                }
            }
        }

        r
    }
}
// @lc end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (vec![vec![1, 0, 0], vec![0, 0, 1], vec![1, 0, 0]], 1),
            (vec![vec![1, 0, 0], vec![0, 1, 0], vec![0, 0, 1]], 3),
        ];
        for (i, (mat, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::num_special(mat);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
