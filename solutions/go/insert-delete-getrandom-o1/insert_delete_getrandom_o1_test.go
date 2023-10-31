package insert_delete_getrandom_o1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_randomizedSet_1(t *testing.T) {
	//	Input
	//["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
	//	[[], [1], [2], [2], [], [1], [2], []]
	//	Output
	//	[null, true, false, true, 2, true, false, 2]

	s := Constructor()
	assert.Equal(t, true, s.Insert(1))
	assert.Equal(t, false, s.Remove(2))
	assert.Equal(t, true, s.Insert(2))
	x := s.GetRandom()
	assert.True(t, x == 1 || x == 2)
	assert.True(t, s.Remove(1))
	assert.False(t, s.Insert(2))
	x = s.GetRandom()
	assert.True(t, x == 2)
}
