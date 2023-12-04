package wordbreak

// Solution 1. DFS + Trie
// Solution 2. DP + Hash
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, w := range wordDict {
		m[w] = true
	}
	n := len(s)
	f := make([]bool, n)
	for i := 0; i < n; i++ {
		for k := range m {
			if f[i] {
				break
			}
			j := i - len(k) + 1
			if j == 0 {
				f[i] = m[s[j:i+1]]
				continue
			}
			if j < 0 || f[j-1] == false || m[s[j:i+1]] == false {
				continue
			}
			f[i] = true
		}
	}
	return f[n-1]
}

func wordBreakV1(s string, wordDict []string) bool {
	root := &TrieNode{}
	for _, s := range wordDict {
		AddToTrie(root, s)
	}
	visited := make(map[int]bool)
	return dfs(s, 0, root, visited)
}

func dfs(s string, p int, root *TrieNode, visited map[int]bool) bool {
	if p == len(s) {
		return true
	}
	if visited[p] {
		return false
	}
	cur := root
	for i := p; i < len(s); i++ {
		v := s[i] - 'a'
		if cur.children[v] == nil {
			break
		}
		cur = cur.children[v]
		if cur.isEnd && dfs(s, i+1, root, visited) {
			return true
		}
	}
	visited[p] = true
	return false
}

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func AddToTrie(root *TrieNode, s string) {
	cur := root
	for i := 0; i < len(s); i++ {
		v := s[i] - 'a'
		if cur.children[v] == nil {
			cur.children[v] = &TrieNode{}
		}
		cur = cur.children[v]
	}
	cur.isEnd = true
}
