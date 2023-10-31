package combination_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	cases := []struct {
		candidates []int
		target     int

		result [][]int
	}{
		{
			[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}},
		},
		{
			[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			[]int{2}, 1, nil,
		},
	}
	for _, c := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, c.result, combinationSum(c.candidates, c.target))
		})
	}
}
