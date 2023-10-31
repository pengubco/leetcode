package largestrectangleinhistogram

import "errors"

/*
f[i]: the max area if heights[i] is the highest bar.
then answer is max{f[i]}.

How do we know heights[i] is the highest bar?
1. The left side is smaller than or equal to it and
2. the right side is smaller than or equal to it.

Maintain an increasing stack. When a bar is smaller than the top of stack, calculate
the area using the top-of-stack as the highest bar.

What if all bars are increasing? Assume there is a bar of 0 height at the end.
*/
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)

	stack := &Stack{}
	n := len(heights)
	ans := 0
	for i := 0; i < n; i++ {
		if stack.IsEmpty() {
			stack.Push(i)
			continue
		}

		for {
			if stack.IsEmpty() {
				break
			}
			top, _ := stack.Top()
			if heights[i] > heights[top] {
				break
			}
			stack.Pop()
			// calculate area using heights[top] as the tallest bar.
			leftBoundary := -1
			if !stack.IsEmpty() {
				leftBoundary, _ = stack.Top()
			}
			area := heights[top] * (i - leftBoundary - 1)
			if ans < area {
				ans = area
			}
		}
		stack.Push(i)
	}
	return ans
}

type Stack struct {
	a    []int
	size int
}

func (s *Stack) Push(x int) {
	s.a = append(s.a, x)
	s.size++
}

func (s *Stack) Top() (int, error) {
	if s.size > 0 {
		return s.a[s.size-1], nil
	}
	return 0, errors.New("empty stack")
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return 0, errors.New("empty stack")
	}
	v := s.a[s.size-1]
	s.size--
	s.a = s.a[:s.size]
	return v, nil
}
