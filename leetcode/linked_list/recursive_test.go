package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	e1 := DeleteDuplicates(&ListNode{1, nil})
	assert.Equal(t, "1", e1.ToString())

	e2 := DeleteDuplicates(&ListNode{1, &ListNode{1, nil}})
	assert.Equal(t, "1", e2.ToString())

	e3 := DeleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{1, nil}}})
	assert.Equal(t, "1", e3.ToString())

	e4 := DeleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}})
	assert.Equal(t, "1", e4.ToString())

	e5 := DeleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, nil}}})
	assert.Equal(t, "1>2", e5.ToString())
}
