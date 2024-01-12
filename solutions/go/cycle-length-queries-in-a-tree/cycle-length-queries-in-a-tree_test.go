package cyclelengthqueriesinatree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycleLengthQueries(t *testing.T) {
	cases := []struct {
		n       int
		queries [][]int
		result  []int
	}{
		{
			3,
			[][]int{{5, 3}, {4, 7}, {2, 3}},
			[]int{4, 5, 3},
		},
		{
			2,
			[][]int{{1, 2}},
			[]int{2},
		},
		{
			5,
			[][]int{{17, 21}, {23, 5}, {15, 7}, {3, 21}, {31, 9}, {5, 15}, {11, 2}, {19, 7}},
			[]int{7, 3, 2, 6, 8, 6, 3, 7},
		},
		{
			4,
			[][]int{{14, 13}, {3, 2}, {14, 3}, {9, 5}, {7, 10}, {12, 4}, {14, 9}, {14, 10}},
			[]int{5, 3, 3, 4, 6, 6, 7, 7},
		},
	}
	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, cycleLengthQueries(tc.n, tc.queries))
		})
	}
}
