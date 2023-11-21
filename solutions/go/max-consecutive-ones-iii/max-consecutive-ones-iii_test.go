package maxconsecutiveonesiii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestOnes(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		result int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, longestOnes(tc.nums, tc.k))
		})
	}
}
