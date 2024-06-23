package longestincreasingsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLIS(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{[]int{1, 2, -10, -8, -7}, 3},
		{[]int{4, 10, 4, 3, 8, 9}, 3},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, lengthOfLIS(tc.nums))
		})
	}
}
