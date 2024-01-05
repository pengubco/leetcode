package numberofoperationstomakenetworkconnected

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeConnected(t *testing.T) {
	cases := []struct {
		n           int
		connections [][]int
		result      int
	}{
		{4, [][]int{{0, 1}, {0, 2}, {1, 2}}, 1},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}, -1},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, makeConnected(tc.n, tc.connections))
		})
	}
}
