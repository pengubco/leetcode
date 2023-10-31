package subarray_sums_divisible_by_k

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subarraysDivByK(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		result int
	}{
		{[]int{4, 5, 0, -2, -3, 1}, 5, 7},
		{[]int{5}, 9, 0},
	}

	for _, c := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, c.result, subarraysDivByK(c.nums, c.k))
		})
	}
}
