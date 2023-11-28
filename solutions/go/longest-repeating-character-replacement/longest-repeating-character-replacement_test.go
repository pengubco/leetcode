package longestrepeatingcharacterreplacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	cases := []struct {
		s      string
		k      int
		result int
	}{
		{"ABCDE", 1, 2},
		{"A", 2, 1},
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"EOEMQLLQTRQDDCOERARHGAAARRBKCCMFTDAQOLOKARBIJBISTGNKBQGKKTALSQNFSABASNOPBMMGDIOETPTDICRBOMBAAHINTFLH", 7, 11},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, characterReplacement(tc.s, tc.k))
		})
	}
}
