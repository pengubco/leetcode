package sumofdistancesintree

func sumOfDistancesInTree(n int, edges [][]int) []int {
	dt := newDistantTree(n, edges)
	dt.dfs(0)
	return dt.distanceSumMap
}

type DistantTree struct {
	n              int
	rootedTree     [][]int
	subtreeSize    []int
	distanceSumMap []int
}

func newDistantTree(n int, edges [][]int) *DistantTree {
	// undirected tree
	ut := make(map[int][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		ut[a] = append(ut[a], b)
		ut[b] = append(ut[b], a)
	}
	t := &DistantTree{
		n:              n,
		rootedTree:     make([][]int, n),
		subtreeSize:    make([]int, n),
		distanceSumMap: make([]int, n),
	}
	visited := make([]bool, n)
	t.initialize(0, ut, visited, 0)
	return t
}

// DFS to initialize
func (t *DistantTree) initialize(cur int, ut map[int][]int, visited []bool, depth int) {
	visited[cur] = true
	t.distanceSumMap[0] += depth
	for _, j := range ut[cur] {
		if visited[j] {
			continue
		}
		t.rootedTree[cur] = append(t.rootedTree[cur], j)
		t.initialize(j, ut, visited, depth+1)
		t.subtreeSize[cur] += t.subtreeSize[j]
	}
	t.subtreeSize[cur] += 1
}

// Pre-order DFS. use parent's distance to calculate a child's distance.
func (t *DistantTree) dfs(cur int) {
	for _, j := range t.rootedTree[cur] {
		t.distanceSumMap[j] = t.distanceSumMap[cur] - t.subtreeSize[j] + (t.n - t.subtreeSize[j])
		t.dfs(j)
	}
}
