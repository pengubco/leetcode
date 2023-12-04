package furthestbuildingyoucanreach

// https://leetcode.com/problems/furthest-building-you-can-reach/

import "container/heap"

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)
	h := make(Heap, 0)

	i := 1
	for ; i < n; i++ {
		d := heights[i] - heights[i-1]
		if d <= 0 {
			continue
		}

		if ladders == 0 {
			if bricks >= d {
				bricks -= d
				continue
			}
			break
		}

		if len(h) < ladders {
			heap.Push(&h, d)
			continue
		}

		// cannot use ladders for all climbs. choose the smallest climb to use bricks.
		if h[0] >= d {
			if bricks >= d {
				bricks -= d
				continue
			} else {
				return i - 1
			}
		}
		v := h[0]
		heap.Pop(&h)
		if bricks >= v {
			bricks -= v
			heap.Push(&h, d)
			continue
		} else {
			return i - 1
		}
	}

	return i - 1
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	n := len(*h)
	ret := (*h)[n-1]
	*h = (*h)[:n-1]
	return ret
}
