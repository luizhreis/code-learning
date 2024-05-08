pub struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let mut start: usize = 0;
        let mut max: usize = 0;
        let mut chars_map: HashMap<char, usize> = HashMap::new();

        for (end, c) in s.chars().enumerate() {
            match chars_map.get(&c) {
                Some(&last_index) => {
                    if start <= last_index {
                        start = last_index + 1
                    }
                }
                None => (),
            };
            chars_map.insert(c, end);

            if (end + 1 - start) > max {
                max = end + 1 - start
            }
        }

        max as i32
    }
}

fn main() {
    Solution::length_of_longest_substring(String::from("abcabcbb"));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_length_of_longest_substring() {
        struct TestCase {
            s: String,
            want: i32
        }

        let mut test_cases: Vec<TestCase> = Vec::new();
        test_cases.push(TestCase {
            s: "abcabcbb".to_string(),
            want: 3
        });
        test_cases.push(TestCase{
            s: "bbbbb".to_string(),
            want: 1
        });
        test_cases.push(TestCase{
            s: "pwwkew".to_string(),
            want: 3
        });

        for test_case in test_cases {
            assert_eq!(Solution::length_of_longest_substring(test_case.s), test_case.want);
        }
    }
}