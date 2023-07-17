package maximalrectangle

import (
	"errors"
)

/*
If we fixed the row of the rectangle's bottom, then the problem is reduced to
the "largest rectangle in histogram".
*/
func maximalRectangle(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])
	heights := make([]int, m)
	ans := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if matrix[r][c] == '0' {
				heights[c] = 0
			} else {
				heights[c] += 1
			}
		}
		// call the rectangle problem
		area := largestRectangleArea(heights)
		if area > ans {
			ans = area
		}
	}
	return ans
}

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
