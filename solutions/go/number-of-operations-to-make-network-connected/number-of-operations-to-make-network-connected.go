package numberofoperationstomakenetworkconnected

// https://leetcode.com/problems/number-of-operations-to-make-network-connected

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	uf := NewUnionFind(n)
	for _, e := range connections {
		uf.Union(e[0], e[1])
	}
	ans := 0
	for i := 1; i < n; i++ {
		if uf.Find(0) != uf.Find(i) {
			uf.Union(0, i)
			ans++
		}
	}
	return ans
}

// UnionFind can union two disjoint sets and find the disjoint set an element belong
// to in O(1) time. Each element is identified by an integer as an index of the
// underlying array.
type UnionFind struct {
	parent []int
	rank   []int

	// satellite data to record. e.g. size of the set, max in the set.
	size []int
}

// NewUnionFind creates a UnionFind that identifies n elements, from 0 to n-1.
func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return &uf
}

// Size returns the number of elements the disjoint set x belongs to.
func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Find(x)]
}

// Find returns the disjoint set x belongs to.
func (uf *UnionFind) Find(x int) int {
	if x != uf.parent[x] {
		uf.parent[x] = uf.Find(uf.parent[x])
		uf.size[x] = uf.size[uf.parent[x]]
	}
	return uf.parent[x]
}

// Union merges the disjoint set x belongs to and the disjoint set y belongs to.
// If x and y belong to the same set, Union is a noop.
func (uf *UnionFind) Union(x, y int) {
	x, y = uf.Find(x), uf.Find(y)
	if x == y {
		return
	}
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
		return
	}
	uf.parent[x] = y
	uf.size[y] += uf.size[x]
	if uf.rank[x] == uf.rank[y] {
		uf.rank[y]++
	}
}
