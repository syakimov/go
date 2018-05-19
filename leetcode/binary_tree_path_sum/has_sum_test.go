package pathsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		Node     *TreeNode
		Sum      int
		Expected bool
	}{
		{
			Node:     NewNode(nil, nil, 1),
			Sum:      1,
			Expected: true,
		},
		{
			Node:     NewNode(NewNode(nil, nil, 1), nil, 1),
			Sum:      2,
			Expected: true,
		},
		{
			Node:     NewNode(NewNode(nil, nil, 2), nil, 1),
			Sum:      2,
			Expected: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.Expected, hasPathSum(tt.Node, tt.Sum))
	}
}
