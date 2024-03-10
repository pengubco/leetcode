package combinationsumii

import (
	"sort"
)

// https://leetcode.com/problems/combination-sum-ii
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	solution := []int{}
	dfs(candidates, target, solution, &result)
	return result
}

func dfs(a []int, target int, solution []int, result *[][]int) {
	if target == 0 {
		newSolution := make([]int, len(solution))
		copy(newSolution, solution)
		*result = append(*result, newSolution)
		return
	}
	for i, v := range a {
		if i > 0 && a[i] == a[i-1] {
			continue
		}
		if target < v {
			break
		}
		solution = append(solution, v)
		dfs(a[i+1:], target-v, solution, result)
		solution = solution[:len(solution)-1]
	}
}

// incorrect, because it cannot handle duplicate. [1,2,5], [1,7], [1,2,5], [1,7]
// the first 1 and second 1 are different 1s from the input array.
func dfs_incorrect(a []int, target int, solution []int, result *[][]int) {
	if target == 0 {
		// should save a new solution
		if len(*result) == 0 || !sameSlice(solution, (*result)[len(*result)-1]) {
			newSolution := make([]int, len(solution))
			copy(newSolution, solution)
			*result = append(*result, newSolution)
		}
		return
	}
	if len(a) == 0 {
		return
	}
	if target < a[0] {
		return
	}
	// pick a[i]
	solution = append(solution, a[0])
	dfs(a[1:], target-a[0], solution, result)
	solution = solution[:len(solution)-1]
	// do not pick a[i]
	dfs(a[1:], target, solution, result)
}

func sameSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
