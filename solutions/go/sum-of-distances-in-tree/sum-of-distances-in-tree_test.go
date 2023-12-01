package sumofdistancesintree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfDistancesInTree(t *testing.T) {

	cases := []struct {
		edges  [][]int
		n      int
		output []int
	}{
		{[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}, 6, []int{8, 12, 6, 10, 10, 10}},
		{[][]int{}, 1, []int{0}},
	}
	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.output, sumOfDistancesInTree(tc.n, tc.edges))
		})
	}
}
