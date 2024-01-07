package maximummatchingofplayerswithtrainers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchPlayersAndTrainers(t *testing.T) {
	cases := []struct {
		players  []int
		trainers []int
		result   int
	}{
		{
			[]int{4, 7, 9},
			[]int{8, 2, 5, 8},
			2,
		},
		{
			[]int{1, 1, 1},
			[]int{10},
			1,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, matchPlayersAndTrainers(tc.players, tc.trainers))
		})
	}
}
