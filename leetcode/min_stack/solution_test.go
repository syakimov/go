package minstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinStack(t *testing.T) {
	s := Constructor()

	s.Push(-2)
	s.Push(0)
	s.Push(-3)

	assert.Equal(t, -3, s.GetMin())
	s.Pop()

	assert.Equal(t, 0, s.Top())

	assert.Equal(t, -2, s.GetMin())
}

func Test_MinStack_failed(t *testing.T) {
	s := Constructor()

	s.Push(-2)
	s.Push(0)
	s.Push(-1)

	assert.Equal(t, -2, s.GetMin())
	assert.Equal(t, -1, s.Top())
	// s.Pop()
	assert.Equal(t, -2, s.GetMin())
}
