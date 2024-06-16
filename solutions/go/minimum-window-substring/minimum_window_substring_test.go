package minimum_window_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minWindow(t *testing.T) {
	cases := []struct {
		s      string
		t      string
		result string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"abc", "cba", "abc"},
	}
	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.result, minWindow(d.s, d.t))
		})
	}
}
