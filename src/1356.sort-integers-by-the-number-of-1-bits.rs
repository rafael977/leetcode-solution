struct Solution;
// @lc start
impl Solution {
    pub fn sort_by_bits(mut arr: Vec<i32>) -> Vec<i32> {
        arr.sort_by(|a, b| {
            let count_a = a.count_ones();
            let count_b = b.count_ones();
            if count_a == count_b {
                return a.cmp(b);
            }
            count_a.cmp(&count_b)
        });
        arr
    }
}
// @lc end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (
                vec![0, 1, 2, 3, 4, 5, 6, 7, 8],
                vec![0, 1, 2, 4, 8, 3, 5, 6, 7],
            ),
            (
                vec![1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1],
                vec![1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024],
            ),
        ];
        for (i, (arr, expected)) in test_cases.iter().enumerate() {
            let got = Solution::sort_by_bits(arr.to_vec());
            let expected = expected.to_vec();
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {:#?}, got: {:#?}",
                i, expected, got
            )
        }
    }
}
