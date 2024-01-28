package stepstomakearraynondecreasing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalSteps(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{
			[]int{10, 6, 5, 10, 15},
			1,
		},
		{
			[]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11},
			3,
		},
		{
			[]int{4, 5, 7, 7, 13},
			0,
		},
		{
			[]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			9,
		},
		{
			[]int{5, 14, 15, 2, 11, 5, 13, 15},
			3,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, totalSteps(tc.nums))
		})
	}
}
