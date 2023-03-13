package combination_sum

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	cases := []struct {
		candidates []int
		target     int

		result [][]int
	}{
		{
			[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}},
		},
		{
			[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			[]int{2}, 1, nil,
		},
	}
	for _, c := range cases {
		t.Run("", func(tt *testing.T) {
			assert.Equal(tt, c.result, combinationSum(c.candidates, c.target))
		})
	}
}

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
