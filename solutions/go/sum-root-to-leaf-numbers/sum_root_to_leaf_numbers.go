package sum_root_to_leaf_numbers

func sumNumbers(root *TreeNode) int {
	var sum int
	dfs(root, 0, &sum)
	return sum
}

func dfs(curr *TreeNode, val int, sum *int) {
	if curr == nil {
		return
	}
	val = val*10 + curr.Val
	if curr.Left == nil && curr.Right == nil {
		*sum += val
		return
	}
	dfs(curr.Left, val, sum)
	dfs(curr.Right, val, sum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
