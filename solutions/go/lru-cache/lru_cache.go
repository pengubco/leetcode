package lru_cache

import "container/list"

/*
Use a double linked list to keep keys in last accessed time. On every operation, move the key to the head. Then the
LRU is tail of the list.
*/

type LRUCache struct {
	cap int
	m   map[int]*list.Element
	l   *list.List
}

type pair struct {
	k, v int
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{}
	c.cap = capacity
	c.m = make(map[int]*list.Element)
	c.l = list.New()
	return c
}

func (c *LRUCache) Get(key int) int {
	res, ok := c.m[key]
	if !ok {
		return -1
	}
	c.l.MoveToFront(res)
	return res.Value.(pair).v
}

func (c *LRUCache) Put(key int, value int) {
	v, ok := c.m[key]
	if ok {
		v.Value = pair{key, value}
		c.l.MoveToFront(v)
		return
	}
	u := c.l.PushFront(pair{key, value})
	c.m[key] = u
	if c.l.Len() > c.cap {
		b := c.l.Back()
		delete(c.m, b.Value.(pair).k)
		c.l.Remove(b)
	}
}
