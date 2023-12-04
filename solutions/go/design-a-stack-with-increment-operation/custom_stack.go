// https://leetcode.com/problems/design-a-stack-with-increment-operation/

/*
key observations.
1. always access the slice from the end.
2. accumulate increments at the right end. carries it to left on pop.
*/
package custom_stack

type CustomStack struct {
	val     []int
	inc     []int
	maxSize int
}

func Constructor(maxSize int) *CustomStack {
	var s CustomStack
	s.val = make([]int, 0, maxSize)
	s.inc = make([]int, 0, maxSize)
	s.maxSize = maxSize
	return &s
}

func (s *CustomStack) Push(x int) {
	if len(s.val) == s.maxSize {
		return
	}
	s.val = append(s.val, x)
	s.inc = append(s.inc, 0)
}

func (s *CustomStack) Pop() int {
	n := len(s.val)
	if n == 0 {
		return -1
	}
	result := s.val[n-1] + s.inc[n-1]
	if n > 1 {
		s.inc[n-2] += s.inc[n-1]
	}
	s.val = s.val[:n-1]
	s.inc = s.inc[:n-1]
	return result
}

func (s *CustomStack) Increment(k int, val int) {
	if k > len(s.val) {
		k = len(s.val)
	}
	if k == 0 {
		return
	}
	s.inc[k-1] += val
}
