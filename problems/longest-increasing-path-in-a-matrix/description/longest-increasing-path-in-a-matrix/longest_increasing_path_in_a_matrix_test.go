package longestincreasingpathinamatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIncreasingPath(t *testing.T) {
	cases := []struct {
		matrix   [][]int
		expected int
	}{
		{[][]int{{0}}, 1},
		{[][]int{{1, 2}}, 2},
		{[][]int{{2, 1}}, 2},
		{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
		{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, 4},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, longestIncreasingPath(tc.matrix))
	}
}
