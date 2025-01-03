package busroutes

// https://leetcode.com/problems/bus-routes

func numBusesToDestination(routes [][]int, source int, target int) int {
	n := len(routes)
	if source == target {
		return 0
	}

	// graph: each node represents a bus route. each bidirectional edge represents
	// two bus routes overlap on some bus stop.
	g := make([][]int, n)

	targetRoutes := make(map[int]struct{})

	// BFS queue and the distance map
	var q []int
	dist := make(map[int]int)

	// build graph
	for i := 0; i < n; i++ {
		s := make(map[int]struct{})
		for _, stop := range routes[i] {
			s[stop] = struct{}{}
			if stop == target {
				targetRoutes[i] = struct{}{}
			}
			if stop == source {
				q = append(q, i)
				dist[i] = 0
			}
		}
		for j := i + 1; j < n; j++ {
			for _, stop := range routes[j] {
				if _, ok := s[stop]; ok {
					g[i] = append(g[i], j)
					g[j] = append(g[j], i)
					break
				}
			}
		}
	}

	if len(targetRoutes) == 0 || len(q) == 0 {
		return -1
	}

	for len(q) > 0 {
		// explore edge u, v
		u := q[0]
		q = q[1:]
		if _, ok := targetRoutes[u]; ok {
			return dist[u] + 1
		}
		for _, v := range g[u] {
			if _, ok := dist[v]; ok {
				continue
			}
			dist[v] = dist[u] + 1
			q = append(q, v)
		}
	}

	return -1
}
