package intervallistintersections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervalIntersection(t *testing.T) {
	cases := []struct {
		firstList  [][]int
		secondList [][]int
		result     [][]int
	}{
		{
			[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			[][]int{{1, 3}, {5, 9}},
			nil,
			nil,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, intervalIntersection(tc.firstList, tc.secondList))
		})
	}
}
