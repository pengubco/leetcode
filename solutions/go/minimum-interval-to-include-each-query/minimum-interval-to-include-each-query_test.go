package minimumintervaltoincludeeachquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInterval(t *testing.T) {
	cases := []struct {
		intervals [][]int
		queries   []int
		result    []int
	}{
		{
			[][]int{
				{1, 4}, {2, 4}, {3, 6}, {4, 4},
			},
			[]int{2, 3, 4, 5},
			[]int{3, 3, 1, 4},
		},
		{
			[][]int{
				{2, 3}, {2, 5}, {1, 8}, {20, 25},
			},
			[]int{2, 19, 5, 22},
			[]int{2, -1, 4, 6},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, minInterval(tc.intervals, tc.queries))
		})
	}
}
