package sumoftwointegers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSum(t *testing.T) {
	assert.Equal(t, 0, getSum(-1, 1))
	assert.Equal(t, 3, getSum(1, 2))
	assert.Equal(t, 3, getSum(2, 1))
	assert.Equal(t, 1, getSum(-2, 3))
	assert.Equal(t, -1, getSum(2, -3))
}
