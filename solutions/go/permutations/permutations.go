package permutations

// https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	dfs(nums, 0, &result)
	return result
}

func dfs(a []int, depth int, result *[][]int) {
	if depth == len(a) {
		tmp := make([]int, depth)
		copy(tmp, a)
		*result = append(*result, tmp)
		return
	}
	for i := depth; i < len(a); i++ {
		a[i], a[depth] = a[depth], a[i]
		dfs(a, depth+1, result)
		a[i], a[depth] = a[depth], a[i]
	}
}
