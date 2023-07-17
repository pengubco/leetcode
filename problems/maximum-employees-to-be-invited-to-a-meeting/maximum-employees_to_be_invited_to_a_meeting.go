package maximumemployeestobeinvitedtoameeting

/*
1. each node has exactly one favorite. a->b. a favorite is b. so there are n directed edges.
2. "sitting at a table" can have two cases.
2.1. for every node, a, its favorite sits at its right side. then the "sitting" is a directed cycle.
2.2. for some nodes, say group A, their favorites sit at their right sides.
for other nodes, say group B, their favorites sit at their left.
Note that, group A and group B can have only one node. We call group A and group B a valid pair.
We can sit as many pairs of group on the table. For example, if there are K couples,
then we can sit all K couples in the table. So the answer for case 2.2 is to sum all
pairs of groups.

case 2.1 is problem 2360. https://leetcode.com/problems/longest-cycle-in-a-graph

case 2.2. calculate length of a group using the same topological sort used to find cycles.
*/
func maximumInvitations(favorite []int) int {
	n := len(favorite)
	// use topological sort to circles.
	// if circle length is 2, then it's case 2.2.
	// if circle length is larger than 3, then it's case 2.1.

	ingrees := make([]int, n)
	depths := make([]int, n)
	for _, v := range favorite {
		ingrees[v]++
	}

	var q []int
	for i := 0; i < n; i++ {
		depths[i] = 1
		if ingrees[i] == 0 {
			q = append(q, i)
		}
	}

	visited := make([]bool, n)
	for len(q) > 0 {
		u := q[0]
		visited[u] = true
		v := favorite[u]
		depths[v] = depths[u] + 1
		ingrees[v]--
		q = q[1:]
		if ingrees[v] == 0 && !visited[v] {
			q = append(q, v)
		}
	}

	// nodes that have not been visited are in circles.

	maxLongCycle := 0
	sumGroupLength := 0

	for v := 0; v < n; v++ {
		if visited[v] {
			continue
		}
		visited[v] = true
		cycleLength := 1
		for cur := favorite[v]; cur != v; cur = favorite[cur] {
			visited[cur] = true
			cycleLength++
		}

		if cycleLength > 2 {
			maxLongCycle = max(maxLongCycle, cycleLength)
		} else {
			sumGroupLength += depths[v] + depths[favorite[v]]
		}
	}

	return max(maxLongCycle, sumGroupLength)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
