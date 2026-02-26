struct Solution;
// @lc start
impl Solution {
    pub fn num_steps(s: String) -> i32 {
        let mut steps = 0;
        let mut bits = s.into_bytes();
        while bits != vec![b'1'] {
            steps += 1;
            if bits.last() == Some(&b'0') {
                bits.pop();
            } else {
                for i in (0..bits.len()).rev() {
                    if bits[i] == b'1' {
                        bits[i] = b'0';
                        if i == 0 {
                            bits.insert(0, b'1');
                            break;
                        }
                    } else {
                        bits[i] = b'1';
                        break;
                    }
                }
            }
        }

        steps
    }
}
// @lc end

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test() {
        let test_cases = vec![
            (String::from("1101"), 6),
            (String::from("10"), 1),
            (String::from("1"), 0),
        ];
        for (i, (s, expected)) in test_cases.into_iter().enumerate() {
            let got = Solution::num_steps(s);
            assert_eq!(
                got, expected,
                "Failed tc{}, wanted: {}, got: {}",
                i, expected, got
            )
        }
    }
}
