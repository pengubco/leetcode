package kthancestorofatreenode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthAncestor(t *testing.T) {
	cases := []struct {
		n        int
		parent   []int
		nodes    []int
		kth      []int
		expected []int
	}{
		// {7, []int{-1, 0, 0, 1, 1, 2, 2}, []int{3, 5, 6}, []int{1, 2, 3}, []int{1, 0, -1}},
		{5, []int{-1, 0, 0, 0, 3}, []int{1, 3, 0, 3, 3}, []int{5, 2, 1, 1, 5}, []int{-1, -1, -1, 0, -1}},
	}

	for _, tc := range cases {
		ta := Constructor(tc.n, tc.parent)
		for i := 0; i < len(tc.nodes); i++ {
			assert.Equal(t, tc.expected[i], ta.GetKthAncestor(tc.nodes[i], tc.kth[i]))
		}
	}
}

/*
              0
				1   2  3
				         4
*/
