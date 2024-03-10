package combinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	cases := []struct {
		n      int
		k      int
		result [][]int
	}{
		{4, 2, [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		}},
		{
			1, 1, [][]int{
				{1},
			},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, combine(tc.n, tc.k))
		})
	}
}
