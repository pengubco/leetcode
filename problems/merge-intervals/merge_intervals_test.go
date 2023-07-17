package mergeintervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		intervals [][]int
		expected  [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}}, [][]int{{1, 4}}},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.expected, merge(tc.intervals))
	}
}
