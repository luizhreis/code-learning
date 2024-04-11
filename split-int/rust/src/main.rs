fn main() {
    let n = 10932;
    println!("n = {}\ndigits = {:?}", n, split_int(n));
}

fn split_int(n: i32) -> Vec<i32> {
    let mut n = n;
    let mut digits = Vec::new();

    while n > 0 {
        digits.push( n % 10 );
        n /= 10;
    }

    digits.reverse();
    digits
}

#[cfg(test)]
mod tests{
    #[test]
    fn it_works(){
        let test_cases = vec![
            (10932, vec![1, 0, 9, 3, 2]),
            (567, vec![5, 6, 7]),
            (0, vec![]),
            (30, vec![3, 0])
        ];

        for (n, expected) in test_cases {
            assert_eq!(crate::split_int(n), expected)
        }
    }
}
