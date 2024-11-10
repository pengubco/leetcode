package singlenumberiii

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		nums   []int
		result []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{[]int{-1, 0}, []int{-1, 0}},
		{[]int{0, 1}, []int{1, 0}},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			sort.Ints(tc.result)
			result := singleNumber(tc.nums)
			sort.Ints(result)
			assert.Equal(t, tc.result, result)
		})
	}
}
