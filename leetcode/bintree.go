package bintree

// TreeNode is a representation of a binary tree node
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
