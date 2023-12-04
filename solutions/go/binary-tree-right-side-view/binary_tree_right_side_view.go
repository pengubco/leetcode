package binary_tree_right_side_view

// https://leetcode.com/problems/binary-tree-right-side-view/

func rightSideView(root *TreeNode) []int {
	mapView := make(map[int]int)
	dfs(root, 0, mapView)
	view := make([]int, len(mapView))
	for k, v := range mapView {
		view[k] = v
	}
	return view
}

func dfs(cur *TreeNode, depth int, view map[int]int) {
	if cur == nil {
		return
	}
	if _, ok := view[depth]; !ok {
		view[depth] = cur.Val
	}
	dfs(cur.Right, depth+1, view)
	dfs(cur.Left, depth+1, view)
}

func rightSideViewBFS(root *TreeNode) []int {
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
