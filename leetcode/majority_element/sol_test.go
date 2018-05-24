package majorityel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		In  []int
		Out int
	}{
		{
			In:  []int{3, 2, 3},
			Out: 3,
		},
		{
			In:  []int{2, 2, 1, 1, 1, 2, 2},
			Out: 2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.Out, majorityElement(tt.In))
	}
}
