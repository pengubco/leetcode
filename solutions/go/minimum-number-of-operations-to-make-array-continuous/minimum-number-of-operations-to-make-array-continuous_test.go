package minimumnumberofoperationstomakearraycontinuous

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{4, 2, 5, 3}, 0},
		{[]int{1, 2, 3, 5, 6}, 1},
		{[]int{1, 10, 100, 1000}, 3},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, minOperations(tc.nums))
		})
	}
}
