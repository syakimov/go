package stock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	assert.Equal(t, maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
	assert.Equal(t, maxProfit([]int{7, 6, 4, 3, 1}), 0)
}
