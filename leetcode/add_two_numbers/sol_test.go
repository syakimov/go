package addtwonums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newNode(v int, n *ListNode) *ListNode {
	return &ListNode{Val: v, Next: n}
}
func Test_addTwoNumbers(t *testing.T) {
	l1 := newNode(2, newNode(4, newNode(3, nil)))
	l2 := newNode(5, newNode(6, newNode(4, nil)))

	r := addTwoNumbers(l1, l2)

	assert.Equal(t, 7, r.Val)
	assert.Equal(t, 0, r.Next.Val)
	assert.Equal(t, 8, r.Next.Next.Val)
}
