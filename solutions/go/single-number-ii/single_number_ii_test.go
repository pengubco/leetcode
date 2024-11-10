package singlenumberii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
		{[]int{-1, -2, -1, -2, -1, -2, 0}, 0},
		{[]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}, -4},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, singleNumber(tc.nums))
		})
	}
}
