package min_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minStack_1(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	assert.Equal(t, -3, s.GetMin())
	s.Pop()
	assert.Equal(t, 0, s.Top())
	assert.Equal(t, -2, s.GetMin())
}

func Test_minStack_2(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	assert.Equal(t, -2, s.GetMin())
	assert.Equal(t, -2, s.Top())
	s.Pop()
}
