package integertoroman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		num      int
		expected string
	}{
		{3, "III"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, intToRoman(tc.num))
		})
	}
}
