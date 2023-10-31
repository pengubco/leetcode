package longestpalindromicsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"bb", "bb"},
		{"ccc", "ccc"},
	}

	for _, tc := range cases {
		assert.Equal(t, len(tc.expected), len(longestPalindrome(tc.s)))
	}
}
