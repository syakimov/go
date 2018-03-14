package tree

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return areEqual(root.Left, root.Right)
}

func areEqual(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val &&
		areEqual(left.Left, right.Right) &&
		areEqual(left.Right, right.Left)
}
