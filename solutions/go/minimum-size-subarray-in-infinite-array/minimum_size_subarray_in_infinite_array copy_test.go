package minimumsizesubarrayininfinitearray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSizeSubarray(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		answer int
	}{
		{[]int{5, 5, 4, 1, 2, 2, 2, 3, 2, 4, 2, 5}, 56, 16},
		{[]int{3, 2, 1, 3, 2, 1, 3, 1, 1, 1, 2, 1, 2, 1, 2, 3, 3, 1}, 78, 41},
		{[]int{2, 1, 5, 7, 7, 1, 6, 3}, 39, 9},
		{[]int{3, 1}, 19, 9},
		{[]int{1, 2, 3}, 5, 2},
		{[]int{1, 1, 12, 3}, 4, 2},
		{[]int{1}, 1, 1},
		{[]int{2}, 1, -1},
		{[]int{2, 4, 6, 8}, 3, -1},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.answer, minSizeSubarray(tc.nums, tc.target))
		})
	}
}
