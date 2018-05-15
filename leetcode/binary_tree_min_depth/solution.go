package binarymin

// TreeNode is a representation of a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _minDepth(root *TreeNode) int {
	return MinDepth(root, 0)
}

// MinDepth should be a private function, Leetcode has messed up here
func MinDepth(root *TreeNode, c int) int {
	if root == nil {
		return c
	}

	if root.Right == nil && root.Left == nil {
		return c + 1
	}

	if root.Left == nil {
		return MinDepth(root.Right, c+1)
	}

	if root.Right == nil {
		return MinDepth(root.Left, c+1)
	}

	lDep := MinDepth(root.Left, c+1)
	rDep := MinDepth(root.Right, c+1)

	if lDep < rDep {
		return lDep
	}

	return rDep
}

// Stolen solution form Leetcode
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
