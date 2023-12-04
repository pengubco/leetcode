package lowest_common_ancestor_of_a_binary_tree

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, lcaNode := dfs(root, p, q)
	return lcaNode
}

func dfs(cur, p, q *TreeNode) (int, *TreeNode) {
	if cur == nil {
		return 0, nil
	}
	cnt1, node1 := dfs(cur.Left, p, q)
	if cnt1 == 2 {
		return 2, node1
	}
	cnt2, node2 := dfs(cur.Right, p, q)
	if cnt2 == 2 {
		return 2, node2
	}
	cnt := cnt1 + cnt2
	if cur == p || cur == q {
		cnt++
	}
	if cnt == 2 {
		return 2, cur
	}
	return cnt, nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
