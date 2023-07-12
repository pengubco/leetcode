package lowest_common_ancestor_of_a_binary_tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, _, lcaNode := dfs(root, p, q)
	return lcaNode
}

// returns a triple.
// is P in the subtree rooted at curr
// is Q in the subtree rooted at curr
// what is the LCA of p and q. nil if LCA is not in the subtree rooted at curr.
func dfs(curr, p, q *TreeNode) (bool, bool, *TreeNode) {
	leftP, leftQ := false, false
	var leftLCA *TreeNode
	if curr.Left != nil {
		leftP, leftQ, leftLCA = dfs(curr.Left, p, q)
	}
	rightP, rightQ := false, false
	var rightLCA *TreeNode
	if curr.Right != nil {
		rightP, rightQ, rightLCA = dfs(curr.Right, p, q)
	}
	if leftLCA != nil {
		return true, true, leftLCA
	}
	if rightLCA != nil {
		return true, true, rightLCA
	}
	foundP := leftP || rightP || curr == p
	foundQ := leftQ || rightQ || curr == q
	var lcaNode *TreeNode
	if foundP && foundQ {
		lcaNode = curr
	}
	return foundP, foundQ, lcaNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
