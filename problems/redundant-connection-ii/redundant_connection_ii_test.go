package redundantconnectionii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantDirectedConnection(t *testing.T) {
	cases := []struct {
		edges    [][]int
		expected []int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}, []int{4, 1}},
		{[][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},
		{[][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}, []int{2, 1}},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, findRedundantDirectedConnection(tc.edges))
	}
}
