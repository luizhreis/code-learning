use std::collections::HashMap;

struct Solution {
    nums: Vec<i32>,
    target:  i32,
}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut numbers: HashMap<i32, i32> = HashMap::new();

        for (i, &n) in nums.iter().enumerate() {
            let dif = target - n;
            
            match numbers.get(&dif) {
                Some(&j) => return vec![j as i32, i as i32],
                None => numbers.insert(n, i as i32),
            };
        }
        unreachable!();
    }
}

fn main() {
    let result = Solution::two_sum(vec![2,7,11,15],9);
    print!("{:?}",result);
}