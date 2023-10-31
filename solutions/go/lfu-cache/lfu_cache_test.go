package lfu_cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LFUCache_1(t *testing.T) {
	/*
		["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
		[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
		Output
		[null, null, null, 1, null, -1, 3, null, -1, 3, 4]
	*/
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	assert.Equal(t, 1, c.Get(1))
	c.Put(3, 3)
	assert.Equal(t, -1, c.Get(2))
	assert.Equal(t, 3, c.Get(3))
	c.Put(4, 4)
	assert.Equal(t, -1, c.Get(1))
	assert.Equal(t, 3, c.Get(3))
	assert.Equal(t, 4, c.Get(4))
}

func Test_LFUCache_2(t *testing.T) {
	/*
		["LFUCache", "put", "put", "get", "get", "get", "put", "get", "get"]
		[[1], [1, 1], [2, 2], [1], [2], [3,3], [2], [3]]
		Output
		[null, null,  null,   -1, 2, 2, null, -1, 3]
	*/
	c := Constructor(1)
	c.Put(1, 1)
	c.Put(2, 2)
	assert.Equal(t, -1, c.Get(1))
	assert.Equal(t, 2, c.Get(2))
	assert.Equal(t, 2, c.Get(2))
	c.Put(3, 3)
	assert.Equal(t, -1, c.Get(2))
	assert.Equal(t, 3, c.Get(3))
}

func Test_LFUCache_3(t *testing.T) {
	/*
	   ["LFUCache","get","put","get","put","put","get","get"]
	   [[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
	   Output:
	   [null, -1, null, -1, null, null, 2, 6]
	*/

	c := Constructor(2)
	assert.Equal(t, -1, c.Get(2))
	c.Put(2, 6)
	assert.Equal(t, -1, c.Get(1))
	c.Put(1, 5)
	c.Put(1, 2)
	assert.Equal(t, 2, c.Get(1))
	assert.Equal(t, 6, c.Get(2))
}

func Test_LFUCache_4(t *testing.T) {
	/*
	   ["LFUCache","put","put","get","get","get","put","put","get","get","get","get"]
	   [[3],[2,2],[1,1],[2],[1],[2],[3,3],[4,4],[3],[2],[1],[4]]
	   Output:
	   [null,null,null,2,1,2,null,null,-1,2,1,4]
	*/
	c := Constructor(3)
	c.Put(2, 2)
	c.Put(1, 1)
	assert.Equal(t, 2, c.Get(2))
	assert.Equal(t, 1, c.Get(1))
	assert.Equal(t, 2, c.Get(2))
	c.Put(3, 3)
	c.Put(4, 4)
	assert.Equal(t, -1, c.Get(3))
	assert.Equal(t, 2, c.Get(2))
	assert.Equal(t, 1, c.Get(1))
	assert.Equal(t, 4, c.Get(4))

}
