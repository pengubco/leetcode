package continuoussubarraysum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSubarraySum(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{23, 2, 4, 6, 7}, 6, true},
		{[]int{23, 2, 6, 4, 7}, 6, true},
		{[]int{23, 2, 6, 4, 7}, 13, false},
	}
	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, checkSubarraySum(tc.nums, tc.k))
		})
	}
}
