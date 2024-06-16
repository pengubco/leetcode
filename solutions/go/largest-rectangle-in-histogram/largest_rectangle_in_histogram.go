package largestrectangleinhistogram

// https://leetcode.com/problems/largest-rectangle-in-histogram/

// https://peng.fyi/post/monotonicity-stack/

func largestRectangleArea(heights []int) int {
	l, r := leftBoundary((heights)), rightBoundary(heights)
	ans := 0
	for i, h := range heights {
		ans = max(ans, h*(r[i]-l[i]+1))
	}
	return ans
}

func leftBoundary(a []int) []int {
	n := len(a)
	l := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		for {
			if len(stack) == 0 {
				break
			}
			if a[stack[len(stack)-1]] < a[i] {
				break
			}
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			l[i] = 0
		} else {
			l[i] = stack[len(stack)-1] + 1
		}
		stack = append(stack, i)
	}
	return l
}

func rightBoundary(a []int) []int {
	n := len(a)
	r := make([]int, n)
	var stack []int
	for i := n - 1; i >= 0; i-- {
		for {
			if len(stack) == 0 {
				break
			}
			if a[stack[len(stack)-1]] < a[i] {
				break
			}
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			r[i] = n - 1
		} else {
			r[i] = stack[len(stack)-1] - 1
		}
		stack = append(stack, i)
	}
	return r
}
