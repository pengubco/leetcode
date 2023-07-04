package word_search_ii

/*
DFS at each cell. At each step, check whether the path is a prefix of all words.
There are at most (3 * 10^4 * 10) prefixes, so we can put all prefixes in hash and check from there.

If the number of prefixes are too many for hashmap, we can use Trie.
*/
func findWords(board [][]byte, words []string) []string {
	// word map. key: prefixes. value: 1 for prefix, 2 for whole word.
	prefixMap := make(map[string]int8)
	for _, s := range words {
		for i := 1; i < len(s); i++ {
			// handles cases where "ab" and "abc" both exist in words
			if prefixMap[s[:i]] != 2 {
				prefixMap[s[:i]] = 1
			}
		}
		prefixMap[s] = 2
	}
	wordsFound := make(map[string]struct{})
	n, m := len(board), len(board[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	resetVisited := func() {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				visited[i][j] = false
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			resetVisited()
			visited[i][j] = true
			dfs(i, j, []byte{board[i][j]}, board, visited, prefixMap, wordsFound)
		}
	}

	var results []string
	for k, _ := range wordsFound {
		results = append(results, k)
	}
	return results
}

// r, c: current row and column.
// n, m: number of rows and number of columns
func dfs(r, c int, path []byte, board [][]byte, visited [][]bool, prefixMap map[string]int8, wordsFound map[string]struct{}) {
	s := string(path)
	v := prefixMap[s]
	if v == 0 {
		return
	}
	if v == 2 {
		wordsFound[s] = struct{}{}
	}
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range directions {
		newR, newC := r+d[0], c+d[1]
		if newR < 0 || newR >= len(board) || newC < 0 || newC >= len(board[0]) || visited[newR][newC] {
			continue
		}
		path = append(path, board[newR][newC])
		visited[newR][newC] = true
		dfs(newR, newC, path, board, visited, prefixMap, wordsFound)
		visited[newR][newC] = false
		path = path[:len(path)-1]
	}
}
