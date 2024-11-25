package waterandjugproblem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanMeasureWater(t *testing.T) {
	cases := []struct {
		x      int
		y      int
		target int
		result bool
	}{
		{3, 5, 4, true},
		{2, 6, 5, false},
		{1, 2, 3, true},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, canMeasureWater(tc.x, tc.y, tc.target))
		})
	}
}
