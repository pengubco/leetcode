package word_search_ii

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findWords(t *testing.T) {
	cases := []struct {
		board      [][]byte
		words      []string
		wordsFound []string
	}{
		{
			[][]byte{
				{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'},
			}, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"},
		},
		{
			[][]byte{
				{'a', 'b'}, {'c', 'd'},
			}, []string{"abcd"}, []string{},
		},
		{
			[][]byte{{'o', 'a', 'b', 'n'}, {'o', 't', 'a', 'e'}, {'a', 'h', 'k', 'r'}, {'a', 'f', 'l', 'v'}},
			[]string{"oa", "oaa"},
			[]string{"oa", "oaa"},
		},
		{
			[][]byte{{'a', 'b', 'c', 'e'}, {'x', 'x', 'c', 'd'}, {'x', 'x', 'b', 'a'}},
			[]string{"abc", "abcd"},
			[]string{"abc", "abcd"},
		},
	}
	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			got := findWords(d.board, d.words)
			if len(d.wordsFound) == 0 {
				assert.Equal(tt, 0, len(got))
				return
			}
			sort.Strings(got)
			sort.Strings(d.wordsFound)
			assert.Equal(tt, d.wordsFound, got)
		})
	}
}
