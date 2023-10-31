package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	sum := addTwoNumbers(l1, l2)
	assert.Equal(t, sum.Val, 7)
	sum = sum.Next
	assert.Equal(t, sum.Val, 0)
	sum = sum.Next
	assert.Equal(t, sum.Val, 8)
}
