package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSum([]int{2, 7, 11, 15}, 9))
}
