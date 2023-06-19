package basic_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculate(t *testing.T) {
	cases := []struct {
		expression string
		result     int
	}{
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
		{"- 3 + 10", 7},
	}
	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.result, calculate(d.expression))
		})
	}
}
