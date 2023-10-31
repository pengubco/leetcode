package maximumemployeestobeinvitedtoameeting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumInvitations(t *testing.T) {
	cases := []struct {
		favorite []int
		expected int
	}{
		{[]int{2, 2, 1, 2}, 3},
		{[]int{1, 2, 0}, 3},
		{[]int{3, 0, 1, 4, 1}, 4},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, maximumInvitations(tc.favorite))
	}
}
