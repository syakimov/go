package balancedtree

import (
	"math"
)

// TreeNode is struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewNode creates new TreeNode
func NewNode(left, right *TreeNode, val int) *TreeNode {
	return &TreeNode{
		Left:  left,
		Right: right,
		Val:   val,
	}
}

func isBalanced(node *TreeNode) bool {
	if node == nil {
		return true
	}

	l := height(node.Left)
	r := height(node.Right)

	if math.Abs(float64(l)-float64(r)) <= 1 && isBalanced(node.Left) && isBalanced(node.Right) {
		return true
	}

	return false
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	hl := height(node.Left)
	hr := height(node.Right)

	if hl > hr {
		return hl + 1
	}

	return 1 + hr
}
