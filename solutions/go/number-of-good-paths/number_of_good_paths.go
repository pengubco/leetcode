package numberofgoodpaths

import "sort"

// https://leetcode.com/problems/number-of-good-paths

/*
If there are X nodes with value M in one subtree, and there are Y node with value M in
another subtree. All nodes are smaller than or equal to M.
Then merging the two subtrees will generate (X * Y) good paths.
This is because there is exactly one path between any two nodes in a tree.

If the are X nodes with value M in a tree and M is the largest value, then there
are X*(X+1)/2 good paths.

In order to guarantee M is the largest value, we only add nodes with value no more than M.
*/

func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	// for simplicity, handle edge cases
	if n == 1 {
		return 1
	}

	ans := 0
	// val to list of nodes with the val
	valToNodes := make(map[int][]int)
	for i, v := range vals {
		valToNodes[v] = append(valToNodes[v], i)
	}

	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	uf := NewUnionFind(n)
	keys := make([]int, 0, len(valToNodes))
	for k := range valToNodes {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, val := range keys {
		for _, node := range valToNodes[val] {
			for _, u := range adj[node] {
				if vals[u] <= val {
					uf.Union(u, node)
				}
			}
		}
		// count number of val in each subtrees
		counts := make(map[int]int)
		for _, node := range valToNodes[val] {
			counts[uf.Find(node)]++
		}
		for _, count := range counts {
			ans += count * (count + 1) / 2
		}
	}

	return ans
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
