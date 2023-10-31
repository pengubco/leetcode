package minimum_number_of_arrows_to_burst_balloons

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMinArrowShots(t *testing.T) {
	cases := []struct {
		points [][]int
		result int
	}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
		{[][]int{{1, 2}}, 1},
	}

	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.result, findMinArrowShots(d.points))
		})
	}

}
