package maximumcompatibilityscoresum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCompatibilitySum(t *testing.T) {
	cases := []struct {
		students [][]int
		mentors  [][]int
		expected int
	}{
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}, [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}}, 8},
		{[][]int{{0, 0}, {0, 0}, {0, 0}}, [][]int{{1, 1}, {1, 1}, {1, 1}}, 0},
		{[][]int{{0, 1, 0, 1, 1, 1}, {1, 0, 0, 1, 0, 1}, {1, 0, 1, 1, 0, 0}}, [][]int{{1, 0, 0, 0, 0, 1}, {0, 1, 0, 0, 1, 1}, {0, 1, 0, 0, 1, 1}}, 10},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.expected, maxCompatibilitySum(tc.students, tc.mentors))
	}
}
