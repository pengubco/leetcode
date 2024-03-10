package combinationsumii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum2(t *testing.T) {
	cases := []struct {
		candidates     []int
		target         int
		expectedResult [][]int
	}{
		{
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			[]int{2, 5, 2, 1, 2},
			5,
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			result := combinationSum2(tc.candidates, tc.target)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
