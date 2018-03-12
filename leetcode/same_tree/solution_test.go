package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToArray(t *testing.T) {
	// Base case with no elements
	var n1 *TreeNode
	s1 := make([]int, 0)
	ToArray(n1, &s1)
	assert.Equal(t, "[]", fmt.Sprint(s1))

	n2 := &TreeNode{1, nil, nil}
	s2 := make([]int, 0)
	ToArray(n2, &s2)
	assert.Equal(t, "[1]", fmt.Sprint(s2))

	n3 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	s3 := make([]int, 0)
	ToArray(n3, &s3)
	assert.Equal(t, "[1 2 3]", fmt.Sprint(s3))

	n4 := &TreeNode{2, &TreeNode{1, &TreeNode{1, &TreeNode{1, nil, nil}, nil}, nil}, &TreeNode{3, nil, nil}}
	s4 := make([]int, 0)
	ToArray(n4, &s4)
	assert.Equal(t, "[1 1 1 2 3]", fmt.Sprint(s4))
}

func TestCompareArrays(t *testing.T) {
	x1 := []int{1, 2, 3}
	y1 := []int{1, 2, 3}
	assert.True(t, CompareArrays(&x1, &y1))

	x2 := []int{1, 2, 3}
	y2 := []int{1, 3}
	assert.False(t, CompareArrays(&x2, &y2))
}

func TestIsSameTree(t *testing.T) {
	// Base true case.
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{1, nil, nil}
	assert.True(t, isSameTree(n1, n2))

	// Base false case.
	t1 := &TreeNode{1, nil, nil}
	t2 := &TreeNode{2, nil, nil}
	assert.False(t, isSameTree(t1, t2))

	// A more complex True case
	tn1 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	tn2 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	assert.True(t, isSameTree(tn1, tn2))

	// A more complex False case
	tnode1 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	tnode2 := &TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}
	assert.False(t, isSameTree(tnode1, tnode2))

	// An nil edge case
	// A more complex False case
	node1 := &TreeNode{1, &TreeNode{1, nil, nil}, nil}
	node2 := &TreeNode{1, nil, &TreeNode{1, nil, nil}}
	assert.False(t, isSameTree(node1, node2))
}
