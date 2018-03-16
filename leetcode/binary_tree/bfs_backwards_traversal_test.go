package traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelTraversal(t *testing.T) {
	// n = 0
	a0 := levelOrderBottom(nil)
	assert.Equal(t, make([][]int, 0), a0)

	// n = 1
	a1 := levelOrderBottom(&TreeNode{1, nil, nil})
	e1 := [][]int{[]int{1}}

	assert.Equal(t, e1, a1)

	// n = 2
	i2 := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	e2 := [][]int{[]int{15, 7}, []int{9, 20}, []int{3}}

	assert.Equal(t, e2, levelOrderBottom(i2))
}
