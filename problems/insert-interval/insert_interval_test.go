package insert_interval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_insert(t *testing.T) {
	cases := []struct {
		intervals   [][]int
		newInterval []int
		result      [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{[][]int{{1, 2}}, []int{3, 4}, [][]int{{1, 2}, {3, 4}}},
		{[][]int{{1, 2}}, []int{2, 3}, [][]int{{1, 3}}},
	}

	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.result, insert(d.intervals, d.newInterval))
		})
	}
}
