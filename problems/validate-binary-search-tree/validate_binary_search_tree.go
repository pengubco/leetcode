package validate_binary_search_tree

import "math"

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt32, math.MaxInt32)
}

func dfs(curr *TreeNode, l, h int) bool {
	if curr == nil {
		return true
	}
	if curr.Val < l || curr.Val > h {
		return false
	}
	return dfs(curr.Left, l, curr.Val-1) && dfs(curr.Right, curr.Val+1, h)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
