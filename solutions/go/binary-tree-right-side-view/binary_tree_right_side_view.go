package binary_tree_right_side_view

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	type pair struct {
		node  *TreeNode
		depth int
	}
	maxDepth := 1
	depth := map[int]*TreeNode{}
	depth[1] = root
	q := []pair{{root, 1}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.node.Left == nil && curr.node.Right == nil {
			continue
		}
		newDepth := curr.depth + 1
		if newDepth > maxDepth {
			maxDepth = newDepth
		}
		if curr.node.Left != nil {
			q = append(q, pair{curr.node.Left, newDepth})
			depth[newDepth] = curr.node.Left
		}
		if curr.node.Right != nil {
			q = append(q, pair{curr.node.Right, newDepth})
			depth[newDepth] = curr.node.Right
		}
	}

	results := make([]int, maxDepth)
	for i := 1; i <= maxDepth; i++ {
		results[i-1] = depth[i].Val
	}
	return results
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
