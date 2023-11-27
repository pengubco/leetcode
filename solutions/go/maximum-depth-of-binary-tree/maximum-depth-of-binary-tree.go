package maximumdepthofbinarytree

func maxDepth(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	return max(maxDepth(cur.Left), maxDepth(cur.Right)) + 1
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
