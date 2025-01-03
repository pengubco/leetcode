package busroutes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumBusesToDestination(t *testing.T) {
	cases := []struct {
		routes [][]int
		source int
		target int
		result int
	}{
		{[][]int{{1}, {1}}, 1, 1, 0},
		{[][]int{{2}, {2, 8}}, 2, 8, 1},
		{[][]int{{1}, {2}}, 1, 2, -1},
		{[][]int{{1}, {2}}, 1, 3, -1},
		{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6, 2},
		{[][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12, -1},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, numBusesToDestination(tc.routes, tc.source, tc.target))
		})
	}
}
