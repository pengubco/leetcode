package balancedbinarytree

func isBalanced(root *TreeNode) bool {
	_, b := dfs(root)
	return b
}

// return the pair (height of tree, is tree balanced).
func dfs(cur *TreeNode) (int, bool) {
	if cur == nil {
		return 0, true
	}
	lHeight, lBalanced := dfs(cur.Left)
	if !lBalanced {
		return -1, false
	}
	rHeight, rBalanced := dfs(cur.Right)
	if !rBalanced {
		return -1, false
	}
	if lHeight-rHeight > 1 || rHeight-lHeight > 1 {
		return -1, false
	}
	return max(lHeight, rHeight) + 1, true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
