package depth

// TreeNode as ususal.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	depth := 0

	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			if q[i] != nil {
				if q[i].Left != nil {
					q = append(q, q[i].Left)
				}
				if q[i].Right != nil {
					q = append(q, q[i].Right)
				}
			}
		}

		q = q[l:]
		depth++
	}

	return depth
}
