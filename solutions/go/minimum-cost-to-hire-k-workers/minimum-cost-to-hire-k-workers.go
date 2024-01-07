package minimumcosttohirekworkers

import (
	"container/heap"
	"errors"
	"math"
	"sort"
)

// https://leetcode.com/problems/minimum-cost-to-hire-k-workers/

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	workers := make([]Worker, n)
	for i := 0; i < n; i++ {
		workers[i] = Worker{
			quality:        quality[i],
			wage:           wage[i],
			wagePerQuality: float64(wage[i]) / float64(quality[i]),
		}
	}
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].wagePerQuality < workers[j].wagePerQuality
	})

	pq, _ := NewPriorityQueue[*Worker](func(v1, v2 *Worker) bool {
		return v1.quality > v2.quality
	})

	minCost := math.MaxFloat64
	totalQualities := 0

	for i := 0; i < n; i++ {
		if pq.Size() < k {
			pq.Push(&workers[i])
			totalQualities += workers[i].quality
			continue
		}
		cost := workers[i-1].wagePerQuality * float64(totalQualities)
		minCost = min(cost, minCost)
		largestQualityWorker, _ := pq.Top()
		if workers[i].quality < largestQualityWorker.quality {
			pq.Pop()
			totalQualities -= largestQualityWorker.quality
			pq.Push(&workers[i])
			totalQualities += workers[i].quality
		}
	}
	cost := workers[n-1].wagePerQuality * float64(totalQualities)
	minCost = min(cost, minCost)

	return minCost
}

type Worker struct {
	quality        int
	wage           int
	wagePerQuality float64
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
