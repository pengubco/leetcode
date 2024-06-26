package maximumcompatibilityscoresum

// https://leetcode.com/problems/maximum-compatibility-score-sum/

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m := len(students)
	scores := make([][]int, m)
	for i := 0; i < m; i++ {
		scores[i] = make([]int, m)
		for j := 0; j < m; j++ {
			scores[i][j] = calculateScore(students[i], mentors[j])
		}
	}
	return dfs(scores, 0, 0)
}

// d: d mentors have matched.
func dfs(scores [][]int, state int, d int) int {
	n := len(scores)
	if d == n {
		return 0
	}
	var ans int
	for i := 0; i < n; i++ {
		if (state & (1 << i)) > 0 {
			continue
		}
		ans = max(ans, scores[i][d]+dfs(scores, state|(1<<i), d+1))
	}
	return ans
}

func calculateScore(a, b []int) int {
	ans := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			ans++
		}
	}
	return ans
}
