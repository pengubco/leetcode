package minimumcosttohirekworkers

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMIncostToHireWorkers(t *testing.T) {
	cases := []struct {
		quality []int
		wage    []int
		k       int
		result  float64
	}{
		{
			[]int{10, 20, 5},
			[]int{70, 50, 30},
			2,
			105.0,
		},
		{
			[]int{3, 1, 10, 10, 1},
			[]int{4, 8, 2, 2, 7},
			3,
			30.66667,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.True(t, math.Abs(tc.result-mincostToHireWorkers(tc.quality, tc.wage, tc.k)) < 0.00001)
		})
	}
}
