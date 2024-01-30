package twobestnonoverlappingevents

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTwoEvents(t *testing.T) {
	cases := []struct {
		events [][]int
		result int
	}{
		{
			[][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}},
			4,
		},
		{
			[][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}},
			5,
		},
		{
			[][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}},
			8,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, maxTwoEvents(tc.events))
		})
	}

}
