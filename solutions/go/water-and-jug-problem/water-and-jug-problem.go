package waterandjugproblem

// https://leetcode.com/problems/water-and-jug-problem/description/

func canMeasureWater(x int, y int, target int) bool {
	if x+y < target {
		return false
	}
	if x+y == target {
		return true
	}
	if x > y {
		x, y = y, x
	}
	p := &Problem{
		x:       x,
		y:       y,
		target:  target,
		visited: make(map[int]struct{}),
	}
	return p.dfs(0, 0)
}

func hash(a, b int) int {
	return a*10000 + b
}

type Problem struct {
	x       int
	y       int
	target  int
	visited map[int]struct{}
}

func (p *Problem) dfs(a, b int) bool {
	if _, ok := p.visited[hash(a, b)]; ok {
		return false
	}
	p.visited[hash(a, b)] = struct{}{}

	if a+b == p.target {
		return true
	}
	// fill jug1
	if p.dfs(p.x, b) {
		return true
	}
	// fill jug2
	if p.dfs(a, p.y) {
		return true
	}
	// empty jug1
	if p.dfs(0, b) {
		return true
	}
	// empty jug2
	if p.dfs(a, 0) {
		return true
	}
	// pour jug1 to jug2
	if a > 0 {
		if p.y-b <= a {
			if p.dfs(a-(p.y-b), p.y) {
				return true
			}
		} else {
			if p.dfs(0, b+a) {
				return true
			}
		}
	}
	// pour jug2 to jug1
	if b > 0 {
		if p.x-a <= b {
			if p.dfs(p.x, b-(p.x-a)) {
				return true
			}
		} else {
			if p.dfs(a+b, 0) {
				return true
			}
		}
	}
	return false
}
