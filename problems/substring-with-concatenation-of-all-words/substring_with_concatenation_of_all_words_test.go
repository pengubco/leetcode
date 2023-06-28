package substring_with_concatenation_of_all_words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	cases := []struct {
		s      string
		words  []string
		result []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
	}

	for _, d := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, d.result, findSubstring(d.s, d.words))
		})
	}
}
