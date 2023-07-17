package flipstringtomonotoneincreasing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFlipsMonoIncr(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{"00110", 1},
		{"010110", 2},
		{"00011000", 2},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, minFlipsMonoIncr(tc.s))
	}
}
