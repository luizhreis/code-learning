#[derive(Debug)]
struct Solution {
    num: String,
    k: i32 
}

impl Solution {
    pub fn remove_kdigits(num: String, mut k: i32) -> String {
        let mut stack = String::new();
        
        for digit in num.chars() {
            while !stack.is_empty() && k > 0 {
                let popped = stack.pop().unwrap();
                if popped > digit {
                    k -= 1;                
                }
                else {
                    stack.push(popped);
                    break;
                }
            }
            stack.push(digit);
        }

        stack.truncate(stack.len() - k as usize);
        stack = stack.trim_start_matches('0' ).to_string();
        if stack.is_empty() {
            "0".to_string()
        }
        else {
            stack
        }
    }
}


fn main() {
    let  solution = Solution{
        num: "1432219".to_string(),
        k: 3
    };
    
    println!("n = {}, got = {}",solution.num, Solution::remove_kdigits("1432219".to_string(),  3));
}

#[cfg(test)]
mod tests{
    #[test]
    fn it_works(){
        let test_cases = vec![
            ("1432219".to_string(), 3, "1219".to_string()),
            ("10200".to_string(), 1, "200".to_string()),
            ("10".to_string(), 2, "0".to_string()),
            ("10".to_string(), 1, "0".to_string()),
            ("26370".to_string(), 3, "20".to_string()),
            ("9".to_string(), 1, "0".to_string()),
            ("112".to_string(), 1, "11".to_string()),
            ("333".to_string(), 1, "33".to_string()),
            ("90001801".to_string(), 2, "101".to_string())
        ];

        for (num, k, want) in test_cases {
            assert_eq!(crate::Solution::remove_kdigits(num, k).to_string(), want.to_string());
        }
    }
}