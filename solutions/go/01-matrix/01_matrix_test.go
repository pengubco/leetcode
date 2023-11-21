package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateMatrix(t *testing.T) {
	cases := []struct {
		mat    [][]int
		result [][]int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
	}
	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, updateMatrix(tc.mat))
		})
	}
}
