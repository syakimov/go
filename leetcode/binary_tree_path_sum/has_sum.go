package pathsum

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

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
