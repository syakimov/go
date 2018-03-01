package main

import "fmt"

// Node...
type Node struct {
	id  int
	lft *Node
	rgt *Node
}

func main() {
	root := &Node{id: 5}
	root.lft = &Node{id: 3}
	root.lft.rgt = &Node{id: 4}
	root.lft.lft = &Node{id: 1}
	root.rgt = &Node{id: 6}
	root.rgt.rgt = &Node{id: 9}

	printTree(root)
}

func printTree(root *Node) {
	if root == nil {
		return
	}

	printTree(root.rgt)
	fmt.Println(root.id)
	printTree(root.lft)
}

func walk(root *Node, f func(n *Node)) {
	if root == nil {
		return
	}

	walk(root.lft, f)

	f(root)

	walk(root.rgt, f)
}

func checkBtree(root *Node, tmp *int) bool {
	if root == nil {
		return true
	}

	checkBtree(root.lft, tmp)

	if tmp == nil {
		tmp = new(int)
		*tmp = root.id
		return true
	}

	ret := root.id > *tmp
	*tmp = root.id

	checkBtree(root.rgt, tmp)

	return ret
}
