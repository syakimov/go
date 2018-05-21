// https://leetcode.com/problems/min-stack/description/

package minstack

// Node keeps the last min value
type Node struct {
	val     int
	lastMin int
}

// MinStack supports get min as well
type MinStack struct {
	nodes []Node
	min   int
}

// Constructor returns instance of MinStack
func Constructor() MinStack {
	return MinStack{nodes: make([]Node, 0), min: -1 << 31}
}

// Push pushes into Minstack
func (this *MinStack) Push(x int) {
	if len(this.nodes) == 0 {
		this.min = x
	}

	this.nodes = append(this.nodes, Node{val: x, lastMin: this.min})
	if this.min > x {
		this.min = x
	}
}

// Pop removes the last added item in the stack and resets min value
func (this *MinStack) Pop() {
	i := len(this.nodes) - 1
	this.min = this.nodes[i].lastMin
	this.nodes = this.nodes[:i]
}

// Top returns the last added item
func (this *MinStack) Top() int {
	return this.nodes[len(this.nodes)-1].val
}

// GetMin returns the min value in the stack
func (this *MinStack) GetMin() int {
	return this.min
}
