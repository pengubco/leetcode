package wordbreak

func wordBreak(s string, wordDict []string) bool {
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
