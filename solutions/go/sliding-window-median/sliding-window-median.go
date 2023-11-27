package slidingwindowmedian

import "container/heap"

func medianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	result := make([]float64, n-k+1)
	if k == 1 {
		for i := 0; i < n; i++ {
			result[i] = float64(nums[i])
		}
		return result
	}

	maxHeap := NewHeapSet[int, int](func(v1, v2 int) bool {
		return v1 > v2
	})
	minHeap := NewHeapSet[int, int](func(v1, v2 int) bool {
		return v1 < v2
	})

	for i := 0; i < n; i++ {
		// delete
		if i >= k {
			if _, ok := maxHeap.Get(i - k); ok {
				maxHeap.Delete(i - k)
			} else {
				minHeap.Delete(i - k)
			}
		}
		// add
		maxHeap.Set(i, nums[i])
		if i < k-1 {
			continue
		}

		// balance to the minHeap
		for maxHeap.Size() > k/2 {
			k, v, _ := maxHeap.Pop()
			minHeap.Set(k, v)
		}
		for {
			k1, v1, _ := maxHeap.Top()
			k2, v2, _ := minHeap.Top()
			if v1 <= v2 {
				break
			}
			maxHeap.Pop()
			minHeap.Pop()
			maxHeap.Set(k2, v2)
			minHeap.Set(k1, v1)
		}

		if k%2 == 0 {
			_, v1, _ := maxHeap.Top()
			_, v2, _ := minHeap.Top()
			result[i-k+1] = float64(v1+v2) / 2
		} else {
			_, v, _ := minHeap.Top()
			result[i-k+1] = float64(v)
		}
	}

	return result
}

// HeapSet keeps key-value pairs in a hash map and provides access to the pair of
// the minimum value.
type HeapSet[K comparable, V any] struct {
	// heap
	h heap.Interface

	// hashmap
	s map[K]*Element[K, V]

	emptyK K
	emptyV V
}

// NewHeapSet returns a HeapSet where values are ordered by the given less function.
func NewHeapSet[K comparable, V any](less func(v1, v2 V) bool) *HeapSet[K, V] {
	hs := HeapSet[K, V]{
		h: newHeapStruct[K, V](less),
		s: make(map[K]*Element[K, V]),
	}
	heap.Init(hs.h)
	return &hs
}

// Set inserts a k-v pair if the key does not exist. Otherwise, Set updates the value.
func (hs *HeapSet[K, V]) Set(k K, v V) {
	existingElement, ok := hs.s[k]
	if !ok {
		e := Element[K, V]{
			Key:   k,
			Value: v,
		}
		heap.Push(hs.h, &e)
		hs.s[k] = &e
		return
	}
	existingElement.Value = v
	heap.Fix(hs.h, existingElement.index)
}

// Get returns the value associated with the key
func (hs *HeapSet[K, V]) Get(k K) (V, bool) {
	e, ok := hs.s[k]
	if !ok {
		return hs.emptyV, false
	}
	return e.Value, true
}

// Delete deletes the key-value pair of the key.
func (hs *HeapSet[K, V]) Delete(key K) {
	item, ok := hs.s[key]
	if !ok {
		return
	}
	delete(hs.s, key)
	index := item.index
	if index == hs.h.Len()-1 {
		hs.h.Pop()
		return
	}
	hs.h.Swap(index, hs.h.Len()-1)
	hs.h.Pop()
	heap.Fix(hs.h, index)
}

// Top returns the key-value pair of the smallest value. It returns false
// if the set is empty.
func (hs *HeapSet[K, V]) Top() (K, V, bool) {
	if hs.h.Len() <= 0 {
		return hs.emptyK, hs.emptyV, false
	}

	var e *Element[K, V]
	h := hs.h.(*heapStruct[K, V])
	e = h.e[0]

	return e.Key, e.Value, true
}

// Pop removes and returns the key-value pair of the smallest value. It returns flase
// if the set is empty.
func (hs *HeapSet[K, V]) Pop() (K, V, bool) {
	if hs.h.Len() == 0 {
		return hs.emptyK, hs.emptyV, false
	}
	e := heap.Pop(hs.h).(*Element[K, V])
	delete(hs.s, e.Key)
	return e.Key, e.Value, true
}

// Size returns the number of key-value pairs.
func (pq *HeapSet[K, V]) Size() int {
	return pq.h.Len()
}

// Map returns the underlying map. It is here to provide an efficient way of
// iterating over all key-value pairs.
func (hs *HeapSet[K, V]) Map() map[K]*Element[K, V] {
	return hs.s
}

// Element is the unit of data stored in hash map and the heap.
type Element[K comparable, V any] struct {
	Key   K
	Value V

	// The array index of the element in the heap.
	index int
}

// heapStruct implements the heap.Interface.
type heapStruct[K comparable, V any] struct {
	e    []*Element[K, V]
	less func(v1, v2 V) bool
}

func newHeapStruct[K comparable, V any](less func(v1, v2 V) bool) *heapStruct[K, V] {
	return &heapStruct[K, V]{
		less: less,
	}
}

func (h *heapStruct[K, V]) Len() int {
	return len(h.e)
}

func (h *heapStruct[K, V]) Less(i, j int) bool {
	return h.less(h.e[i].Value, h.e[j].Value)
}

func (h *heapStruct[K, V]) Swap(i, j int) {
	h.e[i], h.e[j] = h.e[j], h.e[i]
	h.e[i].index = i
	h.e[j].index = j
}

func (h *heapStruct[K, V]) Push(x any) {
	n := len(h.e)
	item := x.(*Element[K, V])
	item.index = n
	h.e = append(h.e, item)
}

func (h *heapStruct[K, V]) Pop() any {
	old := h.e
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	h.e = old[0 : n-1]
	return item
}
