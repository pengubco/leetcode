package redundantconnectionii

/*
The problem guarantees after removing one edge, the digraph is a directed tree.
What this implies.
1. If the graph is undirected, there mush be a circle.
2. Now the graph is directed, so the circle could be
2.1. a directed circle. remove any edge works.
2.2. not a directed circle. Then there must be exactly one node has in degree of 2.
i.e. has two parents. Note that, in a directed tree, each non-root node has 1 indegree,
and the root has 0 indegree. So we are sure that with this extra edge, there is
exactly one node with indegree of 2.
Try remove one of the edge and see whether the undirected version has circles.
If there is no cycle, then it is a valid candidate.
If there is cycle, then it is not valid candidate.
Note that, it's possible that both candidates are valid.
*/

func findRedundantDirectedConnection(edges [][]int) []int {
	// create the adjacent graph
	n := len(edges)
	ingrees := make([][]int, n+1)
	for _, e := range edges {
		a, b := e[0], e[1]
		ingrees[b] = append(ingrees[b], a)
	}

	var candidateEdges [][]int
	for u, parents := range ingrees {
		if len(parents) == 2 {
			candidateEdges = append(candidateEdges, []int{parents[0], u})
			candidateEdges = append(candidateEdges, []int{parents[1], u})
		}
	}

	// case 2.1.
	if candidateEdges == nil {
		return findEdgeToRemove(edges)
	}

	// case 2.2
	// go backwards because the later edge is preferred.
	for i := 1; i >= 0; i-- {
		if !isCircle(edges, candidateEdges[i]) {
			return candidateEdges[i]
		}
	}
	// this return never happens.
	return nil
}

func findEdgeToRemove(edges [][]int) []int {
	n := len(edges)
	uf := NewUnionFind(n + 1)
	for _, e := range edges {
		a, b := e[0], e[1]
		if uf.Find(a) == uf.Find(b) {
			return e
		}
		uf.Union(a, b)
	}
	return nil
}

// Is there a cycle when excludedEdeg is excluded and the directed edge becomes undirected.
func isCircle(edges [][]int, excludedEdge []int) bool {
	n := len(edges)
	uf := NewUnionFind(n + 1)
	for _, e := range edges {
		if e[0] == excludedEdge[0] && e[1] == excludedEdge[1] {
			continue
		}
		a, b := e[0], e[1]
		if uf.Find(a) == uf.Find(b) {
			return true
		}
		uf.Union(a, b)
	}
	return false
}

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return &uf
}

func (uf *UnionFind) Find(x int) int {
	if x != uf.parent[x] {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	x, y = uf.Find(x), uf.Find(y)
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
		return
	}
	uf.parent[x] = y
	if uf.rank[x] == uf.rank[y] {
		uf.rank[y]++
	}
}
