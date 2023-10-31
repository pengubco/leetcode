package wordbreak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	cases := []struct {
		s        string
		words    []string
		expected bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		// {"applepenapple", []string{"apple", "pen"}, true},
		// {"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, wordBreak(tc.s, tc.words))
	}
}
