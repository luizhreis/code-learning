package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	result := &ListNode{Val: 0, Next: nil}
	current := result

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10, Next: nil}
		current = current.Next

	}
	return result.Next
}

func toListNode(nums []int) *ListNode {
	var node *ListNode
	var next *ListNode
	next = nil

	for i := len(nums) - 1; i >= 0; i-- {
		node = &ListNode{Val: nums[i], Next: next}
		next = node
	}
	return node
}
