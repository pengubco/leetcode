package shortestcycleinagraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestCycle(t *testing.T) {
	cases := []struct {
		n        int
		edges    [][]int
		expected int
	}{
		{7, [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3}}, 3},
		{4, [][]int{{0, 1}, {0, 2}}, -1},
		{8, [][]int{{1, 3}, {3, 5}, {5, 7}, {7, 1}, {0, 2}, {2, 4}, {4, 0}, {6, 0}, {6, 1}}, 3},
		{13, [][]int{{0, 1}, {1, 2}, {2, 0}, {0, 3}, {3, 4}, {4, 5}, {6, 7}, {7, 8}, {8, 9}, {9, 10}, {10, 11}, {11, 12}, {12, 0}, {2, 7}, {2, 4}, {1, 8}, {1, 11}}, 3},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, findShortestCycle(tc.n, tc.edges))
	}
}
