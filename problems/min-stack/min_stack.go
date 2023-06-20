package min_stack

/*
Key observation: A stack always pops at the end. So the minimum of a stack is essentially "find the minimum from the first
k elements of an array".
*/

type MinStack struct {
	stack    []int
	minIndex []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	n := len(s.stack)
	if n == 0 {
		s.stack = append(s.stack, val)
		s.minIndex = append(s.minIndex, 0)
		return
	}
	s.stack = append(s.stack, val)
	if val < s.stack[s.minIndex[n-1]] {
		s.minIndex = append(s.minIndex, n)
	} else {
		s.minIndex = append(s.minIndex, s.minIndex[n-1])
	}
}

func (s *MinStack) Pop() {
	n := len(s.stack)
	s.stack = s.stack[:n-1]
	s.minIndex = s.minIndex[:n-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.stack[s.minIndex[len(s.stack)-1]]
}
