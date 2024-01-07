package minimumintervaltoincludeeachquery

// https://leetcode.com/problems/minimum-interval-to-include-each-query/

import (
	"container/heap"
	"errors"
	"sort"
)

func minInterval(intervals [][]int, queries []int) []int {
	ans := make(map[int]int)
	for _, i := range queries {
		ans[i] = -1
	}

	dedupedQueries := SortAndUniq(queries)
	m := len(dedupedQueries)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)

	pq, _ := NewPriorityQueue[[]int](func(v1, v2 []int) bool {
		return v1[1]-v1[0] < v2[1]-v2[0]
	})

	for i, j := 0, 0; j < m; {
		q := dedupedQueries[j]
		if i < n && intervals[i][1] < q {
			i++
			continue
		}
		if (i < n && intervals[i][0] > q) || i >= n {
			// see whether the heap has valid interval
			v := -1
			for pq.Size() > 0 && v == -1 {
				interval, _ := pq.Top()
				if interval[1] < q {
					pq.Pop()
					continue
				}
				v = interval[1] - interval[0] + 1
			}
			ans[q] = v
			j++
			continue
		}
		if i < n {
			pq.Push(intervals[i])
			i++
		}
	}

	result := make([]int, len(queries))
	for i, q := range queries {
		result[i] = ans[q]
	}

	return result
}

func SortAndUniq(a []int) []int {
	m := len(a)
	tmp := make([]int, m)
	copy(tmp, a)
	sort.Ints(tmp)
	deduped := make([]int, 0)
	deduped = append(deduped, tmp[0])
	j := 0
	for i := 1; i < m; i++ {
		if tmp[i] != deduped[j] {
			deduped = append(deduped, tmp[i])
			j++
		}
	}
	return deduped
}

var ErrQueueIsEmpty = errors.New("queue is empty")

// PriorityQueue implements priority queue on any type as long as there is a
// less function to tell the "less-than" relationship between values of the type.
type PriorityQueue[V any] struct {
	hs *heapStruct[V]

	emptyV V
}

// NewPriorityQueue returns a priority queue. Returns error when the less function
// is nil.
func NewPriorityQueue[V any](less func(v1, v2 V) bool) (*PriorityQueue[V], error) {
	if less == nil {
		return nil, errors.New("must provide the compare function")
	}
	pq := &PriorityQueue[V]{
		hs: newHeapStruct[V](less),
	}
	heap.Init(pq.hs)
	return pq, nil
}

// Push inserts a value to the priority queue.
func (pq *PriorityQueue[V]) Push(v V) {
	e := heapElement[V]{
		Value: v,
	}
	heap.Push(pq.hs, &e)
}

// Pop removes and returns the smallest value if the queue is not empty.
func (pq *PriorityQueue[V]) Pop() (V, error) {
	v, err := pq.Top()
	if err != nil {
		return pq.emptyV, err
	}
	heap.Pop(pq.hs)
	return v, nil
}

// Top returns the smallest value if the queue is not empty.
func (pq *PriorityQueue[V]) Top() (V, error) {
	if pq.hs.Len() == 0 {
		return pq.emptyV, ErrQueueIsEmpty
	}
	return pq.hs.e[0].Value, nil
}

// Size returns the number of values in the queue.
func (pq *PriorityQueue[V]) Size() int {
	return pq.hs.Len()
}

// ==== internal types implementing the heap.Interface ====
// heapElement is the unit of data stored in the heap.
type heapElement[V any] struct {
	Value V

	// The array index of the element in the heap.
	index int
}

// heapStruct implements the heap.Interface.
type heapStruct[V any] struct {
	e    []*heapElement[V]
	less func(v1, v2 V) bool
}

func newHeapStruct[V any](less func(v1, v2 V) bool) *heapStruct[V] {
	return &heapStruct[V]{
		less: less,
	}
}

func (h *heapStruct[V]) Len() int {
	return len(h.e)
}

func (h *heapStruct[V]) Less(i, j int) bool {
	return h.less(h.e[i].Value, h.e[j].Value)
}

func (h *heapStruct[V]) Swap(i, j int) {
	h.e[i], h.e[j] = h.e[j], h.e[i]
	h.e[i].index = i
	h.e[j].index = j
}

func (h *heapStruct[V]) Push(x any) {
	n := len(h.e)
	item := x.(*heapElement[V])
	item.index = n
	h.e = append(h.e, item)
}

func (h *heapStruct[V]) Pop() any {
	old := h.e
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	h.e = old[0 : n-1]
	return item
}
