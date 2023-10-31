package largestrectangleinhistogram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	cases := []struct {
		heights  []int
		expected int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
		{[]int{4, 2, 1}, 4},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, largestRectangleArea(tc.heights))
	}
}
