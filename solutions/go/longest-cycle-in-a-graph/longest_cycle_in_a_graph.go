package longestcycleinagraph

/*
Observations:
1. Each node has no outgoing edge or one outgoing edge. If a node has no outgoing edge,
then the node can be ignored because it cannot be in part of a circle.
2. Every node in the circle have at least one incoming edge. If a nodes has no
incoming edge, then the node can be ignored.
3. We can repeat step 1 and step 2 until we cannot apply them.

Now the remaining graph is a disjoint set of circles. We just need to find the largest circles.
*/

func longestCycle(edges []int) int {
	n := len(edges)
	ingress := make([]int, n)
	visited := make([]bool, n)
	for _, v := range edges {
		if v != -1 {
			ingress[v]++
		}
	}

	var q []int
	for i := 0; i < n; i++ {
		if ingress[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		visited[u] = true
		q = q[1:]
		v := edges[u]
		if v == -1 {
			continue
		}
		ingress[v]--
		if visited[v] {
			continue
		}
		if ingress[v] == 0 {
			q = append(q, v)
		}
	}

	maxCycleLength := -1

	for u := 0; u < n; u++ {
		if visited[u] {
			continue
		}
		cycleLength := 1
		visited[u] = true
		for v := edges[u]; v != u; v = edges[v] {
			visited[v] = true
			cycleLength++
		}
		if cycleLength > maxCycleLength {
			maxCycleLength = cycleLength
		}
	}
	return maxCycleLength
}
