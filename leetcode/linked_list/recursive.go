package linkedList

import (
	"fmt"
)

// ListNode is struct corresponding to node in a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// DeleteDuplicates remove the nodes with equal value of a sorted linked list.
func DeleteDuplicates(head *ListNode) *ListNode {
	return deleteDuplicatesRecursive(head)
}

func deleteDuplicatesRecursive(head *ListNode) *ListNode {
	// Recursion bottom gets the following cases:
	// head is nil in the beginning
	// there is only one element
	if head == nil || head.Next == nil {
		return head
	}

	head.Next = deleteDuplicatesRecursive(head.Next)

	// here comes the logic
	if head.Next != nil && head.Val == head.Next.Val {
		return head.Next
	}

	return head
}

// ToString returns string representation of the node format n -> node: n1>n2>n3
func (node *ListNode) ToString() string {
	p := ""
	for node != nil {
		p += fmt.Sprintf("%d>", node.Val)
		node = node.Next
	}

	return p[:len(p)-1]
}
