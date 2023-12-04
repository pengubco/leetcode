package longestincreasingpathinamatrix

// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

/*
1. no need to be worried about visiting a node twice in a path, because it requires increasing.
*/

var dirs = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

type Position struct {
	r int
	c int
}

func longestIncreasingPath(matrix [][]int) int {
	memo := make(map[Position]int)
	ans := 0
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			length := dfs(memo, matrix, i, j)
			if length > ans {
				ans = length
			}
		}
	}
	return ans
}

// longest path starting from m[r][c]
func dfs(memo map[Position]int, matrix [][]int, r int, c int) int {
	if v, ok := memo[Position{r, c}]; ok {
		return v
	}
	maxLength := 0
	for i := 0; i < 4; i++ {
		newR, newC := r+dirs[i][0], c+dirs[i][1]
		if !(newR >= 0 && newR < len(matrix) && newC >= 0 && newC < len(matrix[0])) {
			continue
		}
		if matrix[newR][newC] <= matrix[r][c] {
			continue
		}
		if v, ok := memo[Position{newR, newC}]; ok {
			if maxLength < v {
				maxLength = v
			}
			continue
		}
		length := dfs(memo, matrix, newR, newC)
		if maxLength < length {
			maxLength = length
		}
	}
	memo[Position{r, c}] = 1 + maxLength
	return 1 + maxLength
}
