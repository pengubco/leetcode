package maximalrectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalRectangle(t *testing.T) {
	cases := []struct {
		matrix   [][]byte
		expected int
	}{
		{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, 6},
		{[][]byte{{'1'}}, 1},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, maximalRectangle(tc.matrix))
	}
}
