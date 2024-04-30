package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type testCase struct {
	l1   *ListNode
	l2   *ListNode
	want *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := buildTestCases()

	for _, testCase := range testCases {
		got := addTwoNumbers(testCase.l1, testCase.l2)
		if !isSameList(got, testCase.want) {
			t.Errorf("addTwoNumbers(%v, %v): want: %v, got: %v",
				testCase.l1, testCase.l2, printList(testCase.want), printList(got))
		}
	}
}

func buildTestCases() []testCase {
	testCases := []testCase{}

	testCases = append(testCases, testCase{
		toListNode([]int{2, 4, 3}),
		toListNode([]int{5, 6, 4}),
		toListNode([]int{7, 0, 8}),
	})

	testCases = append(testCases, testCase{
		toListNode([]int{0}),
		toListNode([]int{0}),
		toListNode([]int{0}),
	})

	testCases = append(testCases, testCase{
		toListNode([]int{9, 9, 9, 9, 9, 9, 9}),
		toListNode([]int{9, 9, 9, 9}),
		toListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
	})

	return testCases
}

func isSameList(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func printList(head *ListNode) string {
	var res strings.Builder
	for head != nil {
		fmt.Println("val: ", head.Val)
		res.WriteString(strconv.Itoa(head.Val))
		res.WriteString(" -> ")
		head = head.Next
	}
	res.WriteString("nil")
	return res.String()
}
