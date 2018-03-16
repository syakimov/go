package traversal

// Binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(n)
func levelOrderBottom(root *TreeNode) [][]int {
	r := make([][]int, 0)
	if root == nil {
		return r
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		// Bottleneck, simular to I/O. Could be avoided with reverse list func.
		r = append([][]int{extractVal(q)}, r...)

		l := len(q)

		for _, e := range q {
			if e.Left != nil {
				q = append(q, e.Left)
			}

			if e.Right != nil {
				q = append(q, e.Right)
			}
		}

		q = q[l:]
	}

	return r
}

func extractVal(nodes []*TreeNode) []int {
	l := make([]int, len(nodes))
	for i, e := range nodes {
		l[i] = e.Val
	}

	return l
}
