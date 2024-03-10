package combinations

// https://leetcode.com/problems/combinations/
func combine(n int, k int) [][]int {
	result := [][]int{}
	solution := []int{}
	dfs(1, n, k, solution, &result)
	return result
}

func dfs(v int, n int, k int, solution []int, result *[][]int) {
	if len(solution) == k {
		newSolution := make([]int, k)
		copy(newSolution, solution)
		*result = append(*result, newSolution)
		return
	}
	if v > n {
		return
	}
	// pick v
	solution = append(solution, v)
	dfs(v+1, n, k, solution, result)
	solution = solution[:len(solution)-1]
	// does not pick v
	dfs(v+1, n, k, solution, result)
}
