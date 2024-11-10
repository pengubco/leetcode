package numberofenclaves

const boundayCell = 2

func numEnclaves(g [][]int) int {
	n := len(g)
	m := len(g[0])

	totalOnes := 0
	boundayCells := 0
	isBoundary := func(i, j int) bool {
		return i == 0 || j == 0 || i == n-1 || j == m-1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if g[i][j] == 0 {
				continue
			}
			totalOnes++
			if g[i][j] == boundayCell {
				continue
			}
			if g[i][j] == 1 && isBoundary(i, j) {
				g[i][j] = boundayCell
				boundayCells += bfs(g, i, j)
				continue
			}
		}
	}
	return totalOnes - boundayCells
}

func bfs(g [][]int, i, j int) int {
	n, m := len(g), len(g[0])
	var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var q []Position
	q = append(q, Position{i, j})
	boundayCells := 0
	for len(q) > 0 {
		r, c := q[0].r, q[0].c
		boundayCells++
		q = q[1:]
		for _, dir := range dirs {
			newR, newC := r+dir[0], c+dir[1]
			if newR < 0 || newR >= n || newC < 0 || newC >= m {
				continue
			}
			if g[newR][newC] == 0 || g[newR][newC] == boundayCell {
				continue
			}
			g[newR][newC] = boundayCell
			q = append(q, Position{newR, newC})
		}
	}
	return boundayCells
}

type Position struct {
	r int
	c int
}
