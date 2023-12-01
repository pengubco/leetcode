package longestturbulentsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTurbulenceSize(t *testing.T) {
	cases := []struct {
		nums []int
		ans  int
	}{
		{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}, 5},
		{[]int{4, 8, 12, 16}, 2},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.ans, maxTurbulenceSize(tc.nums))
		})
	}
}
