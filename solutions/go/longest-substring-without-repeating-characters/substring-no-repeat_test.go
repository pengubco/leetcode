package substring_no_repeat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))

	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))

	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}
