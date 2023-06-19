package trapping_rain_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_trap(t *testing.T) {
	cases := []struct {
		height []int
		result int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{4}, 0},
		{[]int{4, 3}, 0},
	}

	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.result, trap(d.height))
		})
	}
}
