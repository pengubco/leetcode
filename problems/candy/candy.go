package candy

func candy(ratings []int) int {
	g := buildDAG(ratings)
	return topologicalSort(g)
}

// Build a DAG graph. each "children" is a node in the graph, represented by the integer of the array index.
// An direct edge (u,v) indicates the 1). u is v's neighbor and 2). rating of u is smaller than the rating of v.
func buildDAG(ratings []int) map[int][]int {
	n := len(ratings)
	g := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		g[i] = []int{}
	}
	for i := 0; i < n-1; i++ {
		if ratings[i] < ratings[i+1] {
			g[i] = append(g[i], i+1)
		} else if ratings[i] > ratings[i+1] {
			g[i+1] = append(g[i+1], i)
		}
	}
	return g
}

func topologicalSort(g map[int][]int) int {
	n := len(g)
	// number of incoming edges to a node.
	inDegree := make(map[int]int, n)
	for i := 0; i < n; i++ {
		inDegree[i] = 0
	}
	for _, l := range g {
		if len(l) == 0 {
			continue
		}
		for _, v := range l {
			inDegree[v]++
		}
	}

	var zeroDegrees []int
	for i, d := range inDegree {
		if d == 0 {
			zeroDegrees = append(zeroDegrees, i)
		}
	}

	sumLevels := 0
	currLevel := 0
	for len(inDegree) > 0 {
		currLevel++
		var newZeroDegrees []int
		for _, u := range zeroDegrees {
			sumLevels += currLevel
			for _, v := range g[u] {
				inDegree[v]--
				if inDegree[v] == 0 {
					newZeroDegrees = append(newZeroDegrees, v)
				}
			}
			delete(inDegree, u)
		}
		zeroDegrees = newZeroDegrees
	}
	return sumLevels
}
