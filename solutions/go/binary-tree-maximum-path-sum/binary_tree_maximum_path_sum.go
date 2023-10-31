package binary_tree_maximum_path_sum

import "math"

/*
There are three types of nodes for a path: 1). internal node of the path; 2). end node of the path; 3). not in the path.
Define following 3 functions.
1. f(x): the max path within subtree rooted at x that ends at x.
f(x) = max{x.Val, f(x.Left) + x.Val, f(x.Right) + x.Val}

2. g(x): the max path within subtree rooted at x that has x as an internal node.
g(x) = max {x.Val, f(x.Left) + f(x.Right) + x.Val}

Why there is no need of a h(x): the max path within subtree rooted at x but does not contain x?
Because that path must be either f(y) or g(y) where y is a node from the subtree.
*/

func maxPathSum(root *TreeNode) int {
	maxPath := math.MinInt32
	dfs(root, &maxPath)
	return maxPath
}

func dfs(curr *TreeNode, maxPath *int) int {
	if curr == nil {
		return 0
	}
	fLeft := dfs(curr.Left, maxPath)
	fRight := dfs(curr.Right, maxPath)
	fCurr := max(max(curr.Val, fLeft+curr.Val), fRight+curr.Val)
	gCurr := max(curr.Val, fLeft+fRight+curr.Val)
	tmp := max(fCurr, gCurr)
	if tmp > *maxPath {
		*maxPath = tmp
	}
	return fCurr
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
