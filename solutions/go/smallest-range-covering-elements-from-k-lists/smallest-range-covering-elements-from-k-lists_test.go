package smallestrangecoveringelementsfromklists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRange(t *testing.T) {
	cases := []struct {
		nums [][]int
		ans  []int
	}{
		{
			[][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}},
			[]int{20, 24},
		},
		{
			[][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
			[]int{1, 1},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.ans, smallestRange(tc.nums))
		})
	}
}
