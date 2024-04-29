pub struct Solution;

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub fn to_list(vec: Vec<i32>) -> List {
    let mut current = None;
    for &v in vec.iter().rev() {
        let mut node = ListNode::new(v);
        node.next = current;
        current = Some(Box::new(node));
    }
    current
}

type List = Option<Box<ListNode>>;

impl Solution {
    pub fn add_two_numbers(l1: List, l2: List) -> List {
      let mut a = &l1;
      let mut b = &l2;
      let mut carry = 0;
      let mut result = ListNode::new(0);
      let  mut current = &mut result;

      while a.is_some() || b.is_some() || carry != 0 {
        let mut sum = carry;
        if let Some(n) = a {
          sum += n.val;
          a = &n.next;
        }
        if let Some(n) = b {
          sum += n.val;
          b = &n.next;
        }
        carry = sum / 10;
        current.next = Some(Box::new(ListNode::new(sum % 10)));
        current = current.next.as_mut().unwrap();
      }
      result.next
    }
}

fn main() {
    let l1 = to_list(vec![2, 4, 3]);
    let l2 = to_list(vec![5, 6, 4]);
    let want = to_list(vec![7, 0, 8]);
    
    assert_eq!(
      Solution::add_two_numbers(l1, l2),
      want
    );
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add_two_numbers() {
      struct TestCase {
        l1: List,
        l2: List,
        want: List
      }

      let mut test_cases: Vec<TestCase> = Vec::new();
      test_cases.push(TestCase {
        l1: to_list(vec![2, 4, 3]),
        l2: to_list(vec![5, 6, 4]),
        want: to_list(vec![7, 0, 8])
      });
      test_cases.push(TestCase {
        l1: to_list(vec![0]),
        l2: to_list(vec![0]),
        want: to_list(vec![0])
      });
      test_cases.push(TestCase {
        l1: to_list(vec![9,9,9,9,9,9,9]),
        l2: to_list(vec![9,9,9,9]),
        want: to_list(vec![8,9,9,9,0,0,0,1])
      });

      for test_case in test_cases {
        assert_eq!(
          Solution::add_two_numbers(test_case.l1, test_case.l2),
          test_case.want
        )
      }
    }
}
