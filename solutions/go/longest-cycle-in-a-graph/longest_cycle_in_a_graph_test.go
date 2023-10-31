package longestcycleinagraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCycle(t *testing.T) {
	cases := []struct {
		edges    []int
		expected int
	}{
		{[]int{3, 3, 4, 2, 3}, 3},
		{[]int{2, -1, 3, 1}, -1},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, longestCycle(tc.edges))
	}
}
