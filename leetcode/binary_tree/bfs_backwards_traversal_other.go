package traversal

// This solution misses tests, so before refactoring please duplicate the tests.

// BreadFirstSearch this is nice implementation of the BFS algorithm.
// It takes pointer to a two dimensional array and fills it with the nodes.
func BreadFirstSearch(root *TreeNode, data *[][]int, level int) {
	if root == nil {
		return
	}

	if level >= len(*data) {
		innerData := make([]int, 0, 2)
		*data = append(*data, innerData)
	}

	// This looks more complicated than it is. Actually what it does is:
	// append the current node to the current level.
	(*data)[level] = append((*data)[level], root.Val)

	// Switch the next two lines to achieve right to left order.
	// If this was a regular tree we could loop and add all of its nodes.
	BreadFirstSearch(root.Left, data, level+1)
	BreadFirstSearch(root.Right, data, level+1)
}

func ReverseSlice(data [][]int) [][]int {
	ret := make([][]int, 0)
	for i := len(data) - 1; i >= 0; i-- {
		ret = append(ret, data[i])
	}
	return ret
}

// LevelOrderBottom is taken from top golang submissions
func LevelOrderBottom(root *TreeNode) [][]int {
	data := make([][]int, 0, 10)

	BreadFirstSearch(root, &data, 0)

	data = ReverseSlice(data)

	return data
}
