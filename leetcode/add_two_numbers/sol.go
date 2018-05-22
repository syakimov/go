// https://leetcode.com/problems/add-two-numbers/description/
// TODO: Not working!

package addtwonums

import (
	"fmt"
	"math"
)

// ListNode is representation of a single linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	res := int(l1.Number() + l2.Number())

	if res == 0 {
		return &ListNode{Val: 0}
	}

	l := &ListNode{Next: &ListNode{}}

	for res > 0 {
		l = l.Next
		l.Val = res % 10
		l.Next = &ListNode{}
		// fmt.Println(l)
		res = res / 10
	}

	fmt.Println(l)
	fmt.Println(l.Next)
	fmt.Println(l.Next.Next)

	return l

}

// Number is good
func (l *ListNode) Number() float64 {
	var c, res float64
	for l != nil {
		res += float64(l.Val) * math.Pow(10, c)
		l = l.Next
		c++
	}

	return res
}
