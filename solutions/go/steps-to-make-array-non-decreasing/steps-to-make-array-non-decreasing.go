package stepstomakearraynondecreasing

import "container/list"

// https://leetcode.com/problems/steps-to-make-array-non-decreasing/
func totalSteps(a []int) int {
	n := len(a)
	f := make([]int, n)
	stack := []int{}
	maxStep := 0
	for i, v := range a {
		step := 0
		// pop from stack if needed, to make sure a[i] can be appended to the decreasing stack.
		for {
			m := len(stack)
			if m == 0 || v < a[stack[m-1]] {
				break
			}
			step = max(step, f[stack[m-1]]+1)
			stack = stack[:m-1]
		}

		if len(stack) > 0 {
			f[i] = max(1, step) // a[i] needs to be removed, so f[i] is at least 1.
			maxStep = max(maxStep, f[i])
		} else {
			f[i] = 0 // a[i] does not need to be removed because there is no larger integer on its left side.
		}

		// push to stack
		stack = append(stack, i)
	}
	return maxStep
}

func totalStepsTLE(nums []int) int {
	l := list.New()
	for _, v := range nums {
		l.PushBack(v)
	}
	rounds := 0
	nodesToDelete := []*list.Element{}
	for {
		for curElement := l.Front(); curElement != nil; curElement = curElement.Next() {
			if curElement.Prev() == nil {
				continue
			}
			if curElement.Prev().Value.(int) > curElement.Value.(int) {
				nodesToDelete = append(nodesToDelete, curElement)
			}
		}
		if len(nodesToDelete) == 0 {
			break
		}
		for _, e := range nodesToDelete {
			l.Remove(e)
		}
		nodesToDelete = nodesToDelete[:0]
		rounds++
	}
	return rounds
}
