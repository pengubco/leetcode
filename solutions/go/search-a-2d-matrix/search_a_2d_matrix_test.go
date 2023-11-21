package searcha2dmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	cases := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true,
		},
		{
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, searchMatrix(tc.matrix, tc.target))
		})
	}
}
