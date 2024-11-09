package furthestbuildingyoucanreach

import "container/heap"

// https://leetcode.com/problems/furthest-building-you-can-reach/

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)
	diffHeap := make(IntHeap, 0)
	i := 1
	for i < n {
		diff := heights[i] - heights[i-1]
		if diff <= 0 {
			i++
			continue
		}
		// use bricks as much as possible
		if bricks >= diff {
			bricks -= diff
			heap.Push(&diffHeap, diff)
			i++
			continue
		}
		// we have to use a ladder
		if ladders == 0 {
			break
		}
		// use the ladder for the largest diff
		if len(diffHeap) > 0 && diffHeap[0] > diff {
			bricks += diffHeap[0]
			heap.Pop(&diffHeap)
			ladders--
		} else {
			ladders--
			i++
		}
	}
	return i - 1
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// We want a max heap
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	n := len(*h)
	ret := (*h)[n-1]
	*h = (*h)[:n-1]
	return ret
}
