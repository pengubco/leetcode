package simplify_path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	cases := []struct {
		path   string
		result string
	}{
		{"/", "/"},
		{"///", "/"},
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
	}
	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.result, simplifyPath(d.path))
		})
	}
}
