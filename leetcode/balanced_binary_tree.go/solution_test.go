package balancedtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalancedBinaryTree(t *testing.T) {
	// fmt.Println(n)
	n := NewNode(nil, nil, 42)
	assert.Equal(t, height(n), 1)
}

func TestIsBalanced(t *testing.T) {
	n := NewNode(
		NewNode(nil, nil, 9),
		NewNode(
			NewNode(nil, nil, 15),
			NewNode(nil, nil, 7),
			20),
		3)

	assert.True(t, isBalanced(n))
}

func printTree(n *TreeNode) {
	if n == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println(n.Val)
	printTree(n.Left)
	printTree(n.Right)
}
