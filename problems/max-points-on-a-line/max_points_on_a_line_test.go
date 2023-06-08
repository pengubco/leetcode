package max_points_on_a_line

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxPoints(t *testing.T) {
	cases := []struct {
		points [][]int
		result int
	}{
		{[][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
		{[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
	}
	for _, c := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, c.result, maxPoints(c.points))
		})
	}
}
