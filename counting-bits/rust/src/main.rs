struct Solution {
    n: i32,
}

impl Solution {
    pub fn count_bits(n: i32) -> Vec<i32> {
        let mut ans = Vec::new();

        for i in 0..=n {
            ans.push( count_ones(i) );
        }
        
        ans
    }
}

pub fn count_ones(mut n: i32) -> i32 {
    let mut count = 0;

    while n > 0 {
        n = n & (n - 1);
        count += 1;
    }

    count
}

fn main() {
    let solution = Solution { n: 5 };

    println!(
        "n = {}, got = {:?}",
        solution.n,
        Solution::count_bits(solution.n)
    );
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        struct TestCase {
            n: i32,
            want: Vec<i32>
        }

        let mut test_cases: Vec<TestCase> = Vec::new();
        test_cases.push(TestCase {n: 0, want: [0].to_vec()});
        test_cases.push(TestCase {n: 2, want: [0, 1, 1].to_vec()});
        test_cases.push(TestCase {n: 5, want: [0, 1, 1, 2, 1, 2].to_vec()});
        test_cases.push(TestCase {n: 8, want: [0, 1, 1, 2, 1, 2, 2, 3, 1].to_vec()});

        for test_case in test_cases {
            assert_eq!(crate::Solution::count_bits(test_case.n), test_case.want)
        }
    }
}


/*
n = 8
0 --> 0000   = 0
1 --> 0001   = 1
2 --> 0010   = 1
3 --> 0011   = 2
4 --> 0100   = 1
5 --> 0101   = 2
6 --> 0110   = 2
7 --> 0111   = 3
8 --> 1000   = 1
*/