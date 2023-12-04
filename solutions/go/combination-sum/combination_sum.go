package combination_sum

import (
	"sort"
)

// https://leetcode.com/problems/combination-sum

func combinationSum(candidates []int, target int) [][]int {
	// Deduplicate and sort candidate in ascending order. It has two advantages:
	// 1. speed up the DFS.
	// 2. avoid duplicate solution.
	// Because the problem guarantees uniqueness, no need to deduplicate.
	sort.Ints(candidates)

	var result [][]int
	var solution []int
	// pass in pointer because Golang always pass-by-value.
	dfs(candidates, target, solution, &result)
	return result
}

func dfs(a []int, target int, solution []int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		newSolution := make([]int, len(solution))
		copy(newSolution, solution)
		*result = append(*result, newSolution)
		return
	}
	for i, v := range a {
		if target < v {
			break
		}
		solution = append(solution, v)
		dfs(a[i:], target-v, solution, result)
		solution = solution[:len(solution)-1]
	}
}
