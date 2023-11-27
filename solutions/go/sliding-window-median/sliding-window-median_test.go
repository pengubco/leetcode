package slidingwindowmedian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianSlidingWindow(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		result []float64
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []float64{1.00000, -1.00000, -1.00000, 3.00000, 5.00000, 6.00000}},
		{[]int{1, 2, 3, 4, 2, 3, 1, 4, 2}, 3, []float64{2.00000, 3.00000, 3.00000, 3.00000, 2.00000, 3.00000, 2.00000}},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, medianSlidingWindow(tc.nums, tc.k))
		})
	}
}
