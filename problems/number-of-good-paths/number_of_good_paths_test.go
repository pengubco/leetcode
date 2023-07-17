package numberofgoodpaths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfGoodPaths(t *testing.T) {
	cases := []struct {
		vals     []int
		edges    [][]int
		expected int
	}{
		{[]int{1, 3, 2, 1, 3}, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, 6},
		{[]int{1, 1, 2, 2, 3}, [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}, 7},
		{[]int{1}, [][]int{}, 1},
		{[]int{1, 1, 1, 1}, [][]int{{0, 1}, {1, 2}, {2, 3}}, 10},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, numberOfGoodPaths(tc.vals, tc.edges))
	}
}
