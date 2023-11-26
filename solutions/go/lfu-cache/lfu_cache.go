package lfu_cache

import "container/list"

// https://leetcode.com/problems/lfu-cache/

type LFUCache struct {
	cap       int
	freqHeads map[int]*list.Element
	kv        map[int]*list.Element
	dll       *list.List
}

type element struct {
	key   int
	value int
	freq  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:       capacity,
		kv:        make(map[int]*list.Element),
		freqHeads: make(map[int]*list.Element),
		dll:       list.New(),
	}
}

func (c *LFUCache) Get(key int) int {
	node, ok := c.kv[key]
	if !ok {
		return -1
	}
	c.increaseFrequencyOfExistingNode(node)
	return node.Value.(*element).value
}

func (c *LFUCache) Put(key int, value int) {
	node, ok := c.kv[key]
	// existing node, no need to evict
	if ok {
		c.increaseFrequencyOfExistingNode(node)
		node.Value.(*element).value = value
		return
	}
	// evict
	if len(c.kv) == c.cap {
		tail := c.dll.Back()
		e := tail.Value.(*element)
		if c.freqHeads[e.freq] == tail {
			delete(c.freqHeads, e.freq)
		}
		delete(c.kv, e.key)
		c.dll.Remove(tail)
	}
	// new node
	c.dll.PushBack(&element{
		key:   key,
		value: value,
		freq:  1,
	})
	node = c.dll.Back()
	c.kv[key] = node
	if c.freqHeads[1] == nil {
		c.freqHeads[1] = node
	} else {
		c.dll.MoveBefore(node, c.freqHeads[1])
		c.freqHeads[1] = node
	}
}

func (c *LFUCache) increaseFrequencyOfExistingNode(node *list.Element) {
	e := node.Value.(*element)

	// case 1: there is only one key.
	if len(c.kv) == 1 {
		delete(c.freqHeads, e.freq)
		c.freqHeads[e.freq+1] = node
		e.freq++
		return
	}
	// case 2: there is only one key with the frequency
	if c.freqHeads[e.freq] == node && (node.Next() == nil || node.Next().Value.(*element).freq != e.freq) {
		delete(c.freqHeads, e.freq)
		if c.freqHeads[e.freq+1] != nil {
			c.dll.MoveBefore(node, c.freqHeads[e.freq+1])
		}
		c.freqHeads[e.freq+1] = node
		e.freq++
		return
	}
	// case 3: there are multiple keys with the frequency, and node is the virtual HEAD
	if c.freqHeads[e.freq] == node && node.Next() != nil && node.Next().Value.(*element).freq == e.freq {
		c.freqHeads[e.freq] = node.Next()
		if c.freqHeads[e.freq+1] != nil {
			c.dll.MoveBefore(node, c.freqHeads[e.freq+1])
		}
		c.freqHeads[e.freq+1] = node
		e.freq++
		return
	}
	// case 4. there are multiple keys with the frequency, and node is not the virtual HEAD
	if c.freqHeads[e.freq] != node {
		if c.freqHeads[e.freq+1] != nil {
			c.dll.MoveBefore(node, c.freqHeads[e.freq+1])
		} else {
			c.dll.MoveBefore(node, c.freqHeads[e.freq])
		}
		c.freqHeads[e.freq+1] = node
		e.freq++
		return
	}
}
