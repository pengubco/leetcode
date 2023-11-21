package matrix

import "math"

func updateMatrix(mat [][]int) [][]int {
	r, c := len(mat), len(mat[0])
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	curRelaxFrom := [][2]int{}
	nextRelaxFrom := [][2]int{}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if mat[i][j] == 0 {
				result[i][j] = 0
				curRelaxFrom = append(curRelaxFrom, [2]int{i, j})
			} else {
				result[i][j] = math.MaxInt
			}
		}
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(curRelaxFrom) > 0 {
		for _, pair := range curRelaxFrom {
			for _, dir := range dirs {
				curRow, curCol := pair[0], pair[1]
				newRow, newCol := curRow+dir[0], curCol+dir[1]
				if newRow < 0 || newRow >= r || newCol < 0 || newCol >= c {
					continue
				}
				if result[curRow][curCol]+1 < result[newRow][newCol] {
					result[newRow][newCol] = result[curRow][curCol] + 1
					nextRelaxFrom = append(nextRelaxFrom, [2]int{newRow, newCol})
				}
			}
		}
		curRelaxFrom, nextRelaxFrom = nextRelaxFrom, curRelaxFrom
		nextRelaxFrom = nextRelaxFrom[:0]
	}
	return result
}
