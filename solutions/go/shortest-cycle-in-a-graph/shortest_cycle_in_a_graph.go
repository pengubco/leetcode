package shortestcycleinagraph

import "math"

// https://leetcode.com/problems/shortest-cycle-in-a-graph/

/*
1. If there are n nodes that have edges, but number of edges is smaller than n, then there is no cycle.
2. How to find the length of a cycle? A cycle is a path plus an edge connecting the start and end of the path.
So we can try each edge (u,v) and calculate the length of the path between u and v, using BFS.
*/
func findShortestCycle(n int, edges [][]int) int {
	m := len(edges)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	nodeWithEdge := 0
	for i := 0; i < n; i++ {
		if len(adj[i]) > 0 {
			nodeWithEdge++
		}
	}
	if m < nodeWithEdge {
		return -1
	}

	shortestPath := math.MaxInt
	for _, e := range edges {
		u, v := e[0], e[1]
		shortestPath = min(shortestPath, bfs(u, v, adj, shortestPath))
	}
	if shortestPath == math.MaxInt {
		return -1
	}
	return shortestPath + 1
}

func bfs(start, end int, adj [][]int, upperbound int) int {
	q := []int{start}
	n := len(adj)
	visited := make([]bool, n)
	distance := make([]int, n)
	distance[start] = 0
	for len(q) > 0 {
		u := q[0]
		visited[u] = true
		q = q[1:]
		for _, v := range adj[u] {
			if visited[v] {
				continue
			}
			// assuming the edge (start,end) does not exist.
			if (u == start && v == end) || (u == end && v == start) {
				continue
			}
			distance[v] = distance[u] + 1
			if v == end {
				return distance[end]
			}
			if distance[v] > upperbound {
				continue
			}
			q = append(q, v)
		}
	}
	if distance[end] == 0 {
		return math.MaxInt
	}
	return distance[end]
}
