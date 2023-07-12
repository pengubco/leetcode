package subarray_sum_equals_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subarraySum(t *testing.T) {
	cases := []struct {
		nums  []int
		k     int
		count int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1}, 0, 0},
	}

	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.count, subarraySum(d.nums, d.k))
		})
	}
}
