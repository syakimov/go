package binarymin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewNode creates new TreeNode
func NewNode(left, right *TreeNode, val int) *TreeNode {
	return &TreeNode{
		Left:  left,
		Right: right,
		Val:   val,
	}
}

func Test_MinDepth(t *testing.T) {
	tests := []struct {
		Input  *TreeNode
		Output int
	}{
		// {
		// Input:  nil,
		// Output: 0,
		// },
		// {
		// Input:  NewNode(nil, nil, 0),
		// Output: 1,
		// },
		{
			Input:  NewNode(NewNode(nil, nil, 0), nil, 0),
			Output: 2,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.Output, minDepth(tt.Input))
	}
}
