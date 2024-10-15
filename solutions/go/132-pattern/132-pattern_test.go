package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind132pattern(t *testing.T) {
	cases := []struct {
		nums   []int
		result bool
	}{
		{[]int{1, -4, 2, -1, 3, -3, -4, 0, -3, -1}, true},
		{[]int{3, 5, 0, 3, 4}, true},
		{[]int{1, 3}, false},
		{[]int{1, 2, 3}, false},
		{[]int{1, 2, 3, 4}, false},
		{[]int{3, 1, 4, 2}, true},
		{[]int{-1, 3, 2, 0}, true},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, find132pattern(tc.nums))
		})
	}
}
