package kthancestorofatreenode

// https://leetcode.com/problems/kth-ancestor-of-a-tree-node/

type TreeAncestor struct {
	log2 []int

	// sparse table
	// st[i][j]: the (2^i)-th parent of node[j]
	st [][]int
}

func Constructor(n int, parent []int) *TreeAncestor {
	log2 := make([]int, n+1)
	log2[1] = 0
	for i := 2; i <= n; i++ {
		log2[i] = log2[i>>1] + 1
	}

	st := make([][]int, log2[n]+1)
	for i := 0; i <= log2[n]; i++ {
		st[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		st[0][i] = parent[i]
	}
	for i := 1; i <= log2[n]; i++ {
		for j := 0; j < n; j++ {
			if st[i-1][j] == -1 {
				st[i][j] = -1
			} else {
				st[i][j] = st[i-1][st[i-1][j]]
			}
		}
	}

	return &TreeAncestor{
		log2: log2,
		st:   st,
	}
}

func (t *TreeAncestor) GetKthAncestor(node int, k int) int {
	for k > 0 {
		i := t.log2[k]
		if t.st[i][node] == -1 {
			return -1
		}
		node = t.st[i][node]
		k -= (1 << i)
	}
	return node
}
