package h_index_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hIndex(t *testing.T) {
	cases := []struct {
		citations []int
		result    int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{0, 0, 0}, 0},
		{[]int{1}, 1},
		{[]int{2}, 1},
		{[]int{0, 10}, 1},
		{[]int{1, 1, 10, 10, 10}, 3},
	}

	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.result, hIndex(d.citations))
		})
	}
}
