package tree

// TreeNode represents node in a tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		if !isSymetricArray(pluckVal(q)) {
			return false
		}

		l := len(q)

		for i := 0; i < l; i++ {
			if q[i] != nil {
				q = append(q, q[i].Left)
				q = append(q, q[i].Right)
			}
		}

		q = q[l:]
	}

	return true
}

//  	1     len 1 -> 0 | 1				1 / 2 = 0
//  1 2 1   len 3 -> 1						3 / 2 = 1
// 1 2 2 1  len 4 -> 2						4 / 2 = 2
func isSymetricArray(a []int) bool {
	for i := 0; i <= len(a)/2; i++ {
		if a[i] != a[len(a)-1-i] {
			return false
		}
	}

	return true
}

func pluckVal(n []*TreeNode) []int {
	r := make([]int, len(n))
	for i, v := range n {
		if v != nil {
			r[i] = v.Val
		} else {
			r[i] = -1
		}

	}

	return r
}

// Some functions I wrote in the mean time
// ======================================
func bfsTreeByOne(root *TreeNode, f func(*TreeNode)) {
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		f(q[0])

		if q[0].Left != nil {
			q = append(q, q[0].Left)
		}

		if q[0].Right != nil {
			q = append(q, q[0].Right)
		}

		q = q[1:]
	}
}

func bfsTreeByWave(root *TreeNode, f func([]*TreeNode)) {
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		l := len(q)
		f(q)

		for i := 0; i < l; i++ {

			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}

			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		q = q[l:]
	}
}
