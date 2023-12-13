package uglynumberii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthUglyNumber(t *testing.T) {
	cases := []struct {
		n      int
		result int
	}{
		{10, 12},
		{1, 1},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, nthUglyNumber(tc.n))
		})
	}
}
