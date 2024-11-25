package escapealargemaze

import "sort"

// https://leetcode.com/problems/escape-a-large-maze/
func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	maxR, maxC := 999999, 999999
	var rIndices, cIndices []int
	for _, loc := range blocked {
		rIndices = append(rIndices, loc[0])
		cIndices = append(cIndices, loc[1])
	}
	// Handle boundaries
	rIndices = append(rIndices, source[0], target[0], 0, maxR)
	cIndices = append(cIndices, source[1], target[1], 0, maxC)

	sort.Ints(rIndices)
	sort.Ints(cIndices)
	rIndexMap, cIndexMap := collapseIndex(rIndices), collapseIndex(cIndices)

	// collapsed grid.
	rowCount, colCount := rIndexMap[maxR]+1, cIndexMap[maxC]+1
	// g[i][j] = true indicates g[i][j] is a blocker.
	g := make([][]bool, rowCount)
	for i := 0; i < rowCount; i++ {
		g[i] = make([]bool, colCount)
	}
	for _, loc := range blocked {
		r, c := rIndexMap[loc[0]], cIndexMap[loc[1]]
		g[r][c] = true
	}

	sourceCell := Cell{rIndexMap[source[0]], cIndexMap[source[1]]}
	targetCell := Cell{rIndexMap[target[0]], cIndexMap[target[1]]}

	// BFS
	queue := []Cell{sourceCell}
	visited := make(map[Cell]struct{})
	visited[sourceCell] = struct{}{}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		cur := queue[0]
		if cur == targetCell {
			return true
		}
		queue = queue[1:]
		for _, dir := range dirs {
			newR, newC := cur.r+dir[0], cur.c+dir[1]
			if newR < 0 || newR >= rowCount || newC < 0 || newC >= colCount {
				continue
			}
			if g[newR][newC] {
				continue
			}
			next := Cell{newR, newC}
			if _, ok := visited[next]; ok {
				continue
			}
			queue = append(queue, next)
			visited[next] = struct{}{}
		}
	}
	return false
}

type Cell struct {
	r int
	c int
}

func collapseIndex(a []int) map[int]int {
	m := make(map[int]int)
	m[a[0]] = 0
	for prev, next := 0, 1; next < len(a); next++ {
		if a[next] == a[next-1] {
			continue
		}
		if a[next]-a[prev] == 1 {
			m[a[next]] = m[a[prev]] + 1
		} else {
			m[a[next]] = m[a[prev]] + 2
		}
		prev = next
	}
	return m
}
