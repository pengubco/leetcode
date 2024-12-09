package smallestrangecoveringelementsfromklists

import (
	"container/heap"
	"errors"
)

// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/description/

func smallestRange(nums [][]int) []int {
	n := len(nums)
	var listNumbers []listNumber
	var ans []int
	listCount := make(map[int]int)
	m := newKWayMerge(nums)
	for {
		ln, err := m.Next()
		if err != nil {
			break
		}
		listNumbers = append(listNumbers, ln)
		listCount[ln.id]++
		for len(listCount) == n {
			newAns := []int{listNumbers[0].value, listNumbers[len(listNumbers)-1].value}
			if ans == nil || newAns[1]-newAns[0] < ans[1]-ans[0] {
				ans = newAns
			}
			ln, listNumbers = listNumbers[0], listNumbers[1:]
			listCount[ln.id]--
			if listCount[ln.id] == 0 {
				delete(listCount, ln.id)
			}
		}
	}

	return ans
}

type kWayMerge struct {
	nums [][]int
	h    *mh
}

func newKWayMerge(nums [][]int) *kWayMerge {
	m := kWayMerge{
		nums: nums,
		h:    &mh{},
	}
	for i, a := range m.nums {
		if len(a) > 0 {
			heap.Push(m.h, listNumber{value: a[0], id: i})
			m.nums[i] = m.nums[i][1:]
		}
	}
	return &m
}

func (m *kWayMerge) Next() (listNumber, error) {
	if m.h.Len() <= 0 {
		return listNumber{}, errors.New("no more number")
	}
	top := heap.Pop(m.h).(listNumber)
	i, a := top.id, m.nums[top.id]
	if len(a) > 0 {
		heap.Push(m.h, listNumber{value: a[0], id: i})
		m.nums[i] = m.nums[i][1:]
	}
	return top, nil
}

type listNumber struct {
	value int
	id    int
}

type mh []listNumber

func (h mh) Len() int {
	return len(h)
}

func (h mh) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h mh) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *mh) Push(x any) {
	*h = append(*h, x.(listNumber))
}

func (h *mh) Pop() any {
	n := len(*h)
	ret := (*h)[n-1]
	*h = (*h)[:n-1]
	return ret
}

/* V1
func smallestRange(nums [][]int) []int {
	listCount := len(nums)

	// For a number, what lists contains it.
	numInList := make(map[int][]int)
	for i, a := range nums {
		for j, x := range a {
			if j == 0 || a[j] != a[j-1] {
				numInList[x] = append(numInList[x], i)
			}
		}
	}

	// min-heap used for merge-sort
	mh := &numberHeap{}
	for i, a := range nums {
		heap.Push(mh, Number{value: a[0], id: i, offset: 0})
	}
	// sorted and unique (de-duped) numbers. using merge-sort.
	a := []int{}
	for mh.Len() > 0 {
		top := heap.Pop(mh).(Number)
		if len(a) == 0 || a[len(a)-1] != top.value {
			a = append(a, top.value)
		}
		if top.offset+1 < len(nums[top.id]) {
			heap.Push(mh, Number{value: nums[top.id][top.offset+1], id: top.id, offset: top.offset + 1})
		}
	}

	var ans []int
	// for each list, how many numbers from the range [ a[i]:a[j] ] are included in the list.
	listContains := make(map[int]int)

	addNumber := func(x int) {
		for _, l := range numInList[x] {
			listContains[l]++
		}
	}

	removeNumber := func(x int) {
		for _, l := range numInList[x] {
			listContains[l]--
			if listContains[l] == 0 {
				delete(listContains, l)
			}
		}
	}

	i, j := 0, -1
	for i < len(a) && j < len(a) {
		// we found an answer
		if len(listContains) == listCount {
			if ans == nil || ans[1]-ans[0] > a[j]-a[i] {
				ans = []int{a[i], a[j]}
			}
			// advance left cursor
			removeNumber(a[i])
			i++
			if i > j {
				j++
				if j < len(a) {
					addNumber(a[j])
				}
			}
			continue
		}
		// advance right cursor
		j++
		if j < len(a) {
			addNumber(a[j])
		}
	}

	return ans
}

type Number struct {
	value int
	// which list the number is from.
	id int
	// what is the offset of the number in the list.
	offset int
}

type numberHeap []Number

// heap.Interface
func (h numberHeap) Len() int {
	return len(h)
}

func (h numberHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h numberHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *numberHeap) Push(x any) {
	*h = append(*h, x.(Number))
}

func (h *numberHeap) Pop() any {
	n := len(*h)
	ret := (*h)[n-1]
	*h = (*h)[:n-1]
	return ret
}
*/
