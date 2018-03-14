package depth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepth(t *testing.T) {
	// n = 0
	r0 := maxDepth(nil)
	assert.Equal(t, 0, r0)

	// n = 1
	r1 := maxDepth(&TreeNode{1, nil, nil})
	assert.Equal(t, 1, r1)

	// n = 2
	r2 := maxDepth(&TreeNode{1, &TreeNode{2, nil, nil}, nil})
	assert.Equal(t, 2, r2)
}
