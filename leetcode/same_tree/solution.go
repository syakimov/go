package tree

// TreeNode is representation of a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}

	return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// ToArray is a function which returns string representation of a node.
func ToArray(n *TreeNode, s *[]int) {
	if n == nil {
		return
	}

	ToArray(n.Left, s)
	*s = append(*s, n.Val)
	ToArray(n.Right, s)
}

// CompareArrays compares two arrays by value. Maybe move somewhere else.
func CompareArrays(a1 *[]int, a2 *[]int) bool {
	if len(*a1) != len(*a2) {
		return false
	}

	for i := 0; i < len(*a1); i++ {
		if (*a1)[i] != (*a2)[i] {
			return false
		}
	}

	return true
}
