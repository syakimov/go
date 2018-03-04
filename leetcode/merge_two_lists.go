// https://leetcode.com/problems/merge-two-sorted-lists/
package main

import "fmt"

func main() {
	// Input: 1->2->4, 1->3->4
	// Output: 1->1->2->3->4->4
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	n := mergeTwoLists(l1, l2)

	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}
