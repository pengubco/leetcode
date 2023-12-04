package redundantconnection

// https://leetcode.com/problems/redundant-connection/

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := NewUnionFind(n + 1)
	for i := 0; i < n; i++ {
		a, b := edges[i][0], edges[i][1]
		if uf.Find(a) == uf.Find(b) {
			return []int{a, b}
		}
		uf.Union(a, b)
	}
	return nil
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
