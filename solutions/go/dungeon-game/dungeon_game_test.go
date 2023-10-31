package dungeon_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateMinimumHP(t *testing.T) {
	cases := []struct {
		points [][]int
		minHP  int
	}{
		{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, 7},
		{[][]int{{0}}, 1},
		{[][]int{{100}}, 1},
	}

	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.minHP, calculateMinimumHP(d.points))
		})
	}
}
