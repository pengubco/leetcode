package furthestbuildingyoucanreach

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFurthestBuilding(t *testing.T) {
	cases := []struct {
		heights  []int
		bricks   int
		ladders  int
		expected int
	}{
		{[]int{4, 2, 7, 6, 9, 14, 12}, 5, 1, 4},
		{[]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2, 7},
		{[]int{14, 3, 19, 3}, 17, 0, 3},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.expected, furthestBuilding(tc.heights, tc.bricks, tc.ladders))
	}
}
