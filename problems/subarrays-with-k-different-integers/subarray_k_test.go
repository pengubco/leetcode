package subarray_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	cases := []struct {
		a        []int
		k        int
		expected int
	}{
		{[]int{1}, 1, 1},
		{[]int{1, 2, 3}, 4, 0},
		{[]int{1, 2}, 1, 2},
		{[]int{1, 2, 1, 2, 3}, 2, 7},
		{[]int{1, 2, 1, 3, 4}, 3, 3},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, subarraysWithKDistinct(tc.a, tc.k))
	}
}
