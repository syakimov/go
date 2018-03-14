package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymetric(t *testing.T) {
	n := &TreeNode{1, nil, nil}
	assert.True(t, isSymmetric(n))

	n1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{2, nil, nil}}
	assert.True(t, isSymmetric(n1))

	n2 := &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}
	assert.False(t, isSymmetric(n2))

	c1 := &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}
	c2 := &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}

	n3 := &TreeNode{1, c1, c2}
	assert.True(t, isSymmetric(n3))

	ch1 := &TreeNode{2, nil, &TreeNode{3, nil, nil}}
	ch2 := &TreeNode{2, nil, &TreeNode{3, nil, nil}}

	n4 := &TreeNode{1, ch1, ch2}
	assert.False(t, isSymmetric(n4))
}
