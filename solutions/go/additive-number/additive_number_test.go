package additivenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAdditiveNumber(t *testing.T) {
	cases := []struct {
		num    string
		result bool
	}{
		{"112358", true},
		{"199100199", true},
		{"1023", false},
		{"101", true},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, isAdditiveNumber(tc.num))
		})
	}
}
