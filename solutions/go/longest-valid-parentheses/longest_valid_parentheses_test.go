package longest_valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	cases := []struct {
		s      string
		result int
	}{
		{"(()", 2},
		{")()())", 4},
		{"", 0},
		{"()", 2},
		{"(()", 2},
	}
	for _, c := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, c.result, longestValidParentheses(c.s))
		})
	}
}
