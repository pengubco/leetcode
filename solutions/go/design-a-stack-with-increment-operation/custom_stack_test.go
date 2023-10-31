package custom_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomStack(t *testing.T) {
	s := Constructor(10)
	assert.Equal(t, -1, s.Pop())

	s.Push(10)
	s.Increment(2, 10)
	s.Push(20)
	s.Increment(10, 10)
	s.Push(10)
	s.Increment(1, 100)
	assert.Equal(t, 10, s.Pop())
	assert.Equal(t, 30, s.Pop())
	assert.Equal(t, 130, s.Pop())
	assert.Equal(t, -1, s.Pop())
}
