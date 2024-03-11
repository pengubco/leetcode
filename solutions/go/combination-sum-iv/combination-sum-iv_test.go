package combinationsumiv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum4(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		result int
	}{
		{[]int{1, 2, 3}, 4, 7},
		{[]int{9}, 3, 0},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, combinationSum4(tc.nums, tc.target))
		})
	}
}
