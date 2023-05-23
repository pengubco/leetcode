package sliding_window_maximum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSlidingWindow(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		result []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{1, 3, -1, -3, 5, 3, 6, 7}},
		{[]int{1}, 1, []int{1}},
	}
	for _, c := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, c.result, maxSlidingWindow(c.nums, c.k))
		})
	}
}
