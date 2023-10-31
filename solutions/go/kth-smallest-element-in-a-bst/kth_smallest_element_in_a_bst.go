package kth_smallest_element_in_a_bst

var order int
var result int

func kthSmallest(root *TreeNode, k int) int {
	order = k
	dfs(root)
	return result
}

func dfs(curr *TreeNode) {
	if curr == nil || order < 0 {
		return
	}
	dfs(curr.Left)
	order--
	if order == 0 {
		result = curr.Val
		return
	}
	dfs(curr.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
