package superuglynumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthSuperUglyNumber(t *testing.T) {
	cases := []struct {
		n      int
		primes []int
		result int
	}{
		{12, []int{2, 7, 13, 19}, 32},
		{1, []int{2, 3, 5}, 1},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, nthSuperUglyNumber(tc.n, tc.primes))
		})
	}
}
