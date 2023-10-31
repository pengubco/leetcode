package three_sum

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_threeSum(t *testing.T) {
	cases := []struct {
		nums   []int
		result [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, nil},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}, [][]int{
			{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1},
		}},
	}
	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			sortIntSlice(d.result)
			result := threeSum(d.nums)
			sortIntSlice(result)
			assert.Equal(tt, d.result, result)
		})
	}
}

func sortIntSlice(a [][]int) {
	sort.Slice(a, func(i, j int) bool {
		n := len(a[i])
		for k := 0; k < n; k++ {
			if a[i][k] != a[j][k] {
				return a[i][k] < a[j][k]
			}
		}
		return false
	})
}
