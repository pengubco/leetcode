package evaluate_reverse_polish_notation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	cases := []struct {
		tokens []string
		result int
	}{
		{[]string{"10"}, 10},
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.result, evalRPN(d.tokens))
		})
	}
}
