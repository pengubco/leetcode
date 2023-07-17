package redundantconnection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantConnection(t *testing.T) {
	cases := []struct {
		edges    [][]int
		expected []int
	}{
		{[][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},

		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, []int{1, 4}},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, findRedundantConnection(tc.edges))
	}
}
